package main

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var globalSession *mgo.Session

func openSession() (session *mgo.Session) {
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		session = nil
		panic(err)
	}
	return
}

/*func closeSession(){
	globalSession.Close()
}*/

func loadSettings() (settings Settings) {
	SettingsCollection := globalSession.DB("db").C("settings")

	query := bson.M{}

	err := SettingsCollection.Find(query).One(&settings)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func changeDBLocale() {
	SettingsCollection := globalSession.DB("db").C("settings")

	err := SettingsCollection.Update(bson.M{"_id": appSettings.Id}, bson.M{"$set": bson.M{"locale": appSettings.Locale}})
	if err != nil {
		fmt.Println(err)
	}
	loadSettings()
}

func getGroups() (groups []Group) {

	GroupsCollection := globalSession.DB("db").C("groups")

	query := bson.M{}

	err := GroupsCollection.Find(query).All(&groups)
	if err != nil {
		fmt.Println(err.Error())
	}

	return
}

func getSubjects() (subjects []Subject) {

	SubjectsCollection := globalSession.DB("db").C("subjects")

	query := bson.M{}

	err := SubjectsCollection.Find(query).All(&subjects)
	if err != nil {
		fmt.Println(err.Error())
	}

	return
}

func getSubject(id string) (subject Subject, err error) {

	SubjectsCollection := globalSession.DB("db").C("subjects")

	// возможно несовпадение типов
	query := bson.M{
		"id": id,
	}

	err = SubjectsCollection.Find(query).One(&subject)
	if err != nil {
		return
	}

	return

}
