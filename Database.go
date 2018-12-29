package MFTCommon

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
)

/**
    CREATE TABLE flashinfo
    (
    uid serial NOT NULL,
    Vendor TEXT,
	Product          TEXT,
	Version          TEXT,
	Title            TEXT,
	Description      TEXT,
    ReleaseDate      TEXT,
	DownloadFileSize TEXT,
	DownloadURL      TEXT,
	DownloadPath     TEXT,
	CurrentPath      TEXT
    )
    WITH (OIDS=FALSE);
*/

type DataBase struct {
	ES  *elastic.Client
	log *logrus.Logger
}

func DBConnect(log *logrus.Logger) DataBase {
	client, err := elastic.NewClient()
	checkErr(err)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		log.WithError(err).Panic("Could not connect to elastic")
	}
	log.Infof("Elasticsearch version %s", esversion)

	return DataBase{ES: client, log: log}
}

func (d DataBase) StoreElement(index string, typeString *string, entry interface{}, id *string) {

	is := d.ES.Index().BodyJson(entry)
	store(d.log, is, index, typeString, id)

}

func (d DataBase) appendElement(index string, typeString *string, entry interface{}, id *string, newElement interface{}) {
	ctx := context.Background()
	scriptString := fmt.Sprintf("ctx._source.%s += newEntry", newElement)
	script := elastic.NewScript(scriptString).
		Params(map[string]interface{}{"newEntry": newElement})

	_, err := d.ES.Update().
		Index(index).
		Type(*typeString).
		Id(*id).
		Script(script).
		Do(ctx)
	if err != nil {
		d.log.Errorf("Error while appending to %s", entry)
	}
	d.log.Infof("Appent %s to %s", newElement, entry)
}

func (d DataBase) updateElement(index string, typeString *string, entry interface{}, id *string) {
	ctx := context.Background()
	scriptString := fmt.Sprintf("ctx._source.%s = newEntry", entry)
	script := elastic.NewScript(scriptString).
		Params(map[string]interface{}{"newEntry": entry})
	_, err := d.ES.Update().
		Index(index).
		Type(*typeString).
		Id(*id).
		Script(script).
		Do(ctx)
	if err != nil {
		d.log.Errorf("Error while updating %s", entry)
	}
	d.log.Infof("updated %s", entry)
}

func (d DataBase) StoreJSON(index string, typeString *string, entry string, id *string) {

	is := d.ES.Index().BodyString(entry)
	store(d.log, is, index, typeString, id)
}

func (d DataBase) Flush(index string) {
	_, err := d.ES.Flush().Index(index).Do(context.Background())
	if err != nil {
		panic(err)
	}
}

func (d DataBase) Search(index string, terms map[string]string) error {
	termQuery := elastic.NewTermQuery("user", "olivere")
	searchResult, err := d.ES.Search().
		Index(index).
		Query(termQuery).
		Sort("user", true).
		From(0).Size(10).
		Pretty(false).
		Do(context.Background())

	if err != nil {
		switch {
		case elastic.IsNotFound(err):
			return err
		case elastic.IsTimeout(err):
			d.log.WithError(err).Error("Timeout retrieving document: %v", err)
			return err
		case elastic.IsConnErr(err):
			d.log.WithError(err).Error("Connection problem: %v", err)
			return err
		default:
			d.log.WithError(err).Error("Unknown error: %v", err)
			return err
		}
	}

	d.log.WithField("query", termQuery).Debug("Query took %d milliseconds\n", searchResult.TookInMillis)
	return nil
}

// Search with a term query

func (d DataBase) Exists(index string, id string) (bool, error, *elastic.GetResult) {
	get, err := d.ES.Get().
		Index(index).
		Id(id).
		Do(context.Background())

	if err != nil {
		switch {
		case elastic.IsNotFound(err):
			return false, nil, get
		case elastic.IsTimeout(err):
			d.log.WithError(err).Errorf("Timeout retrieving document: %v", err)
			return false, err, get
		case elastic.IsConnErr(err):
			d.log.WithError(err).Errorf("Connection problem: %v", err)
			return false, err, get
		default:
			d.log.WithError(err).Errorf("Unknown error: %v", err)
			return false, err, get
		}
	}
	return true, nil, get
}

func store(log *logrus.Logger, is *elastic.IndexService, index string, typeString *string, id *string) error {
	is = is.Index(index)

	if typeString != nil {
		is = is.Type(*typeString)
	} else {
		is.Type(index)
	}

	if id != nil {
		is = is.Id(*id)
	}

	put1, err := is.Do(context.Background())
	if err != nil {
		// Handle error
		log.WithError(err).Error("Could not execute elastic search")
		return err
	}
	log.WithField("dbentry", put1).Infof("Indexed %s to index %s, type %s", put1.Id, put1.Index, put1.Type)
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
