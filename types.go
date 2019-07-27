package main

import "gopkg.in/mgo.v2/bson"

type Settings struct {
	Id     bson.ObjectId `bson:"_id"`
	Locale string        `bson:"locale"`
}

type Data struct {
	Thing  interface{}
	Locale Lang
}

type Group struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
}

type Subject struct {
	Id     bson.ObjectId `bson:"_id"`
	Name   string        `bson:"name"`
	Themes []Theme       `bson:"themes"`
}

type Theme struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
}

type jsonLang struct {
	Name  string `json:"name"`
	Short string `json:"short"`
}

type Lang struct {
	CurrentLocale     string   `json:"locale"`
	System            string   `json:"system"`
	Student           string   `json:"student"`
	Teacher           string   `json:"teacher"`
	StudentOrTeacher  string   `json:"studentOrTeacher"`
	YourIPAddress     string   `json:"yourIPAddress"`
	Subject           string   `json:"subject"`
	Theme             string   `json:"theme"`
	GoToTesting       string   `json:"goToTesting"`
	StartTest         string   `json:"startTest"`
	Back              string   `json:"back"`
	Surname           string   `json:"surname"`
	Name              string   `json:"name"`
	Group             string   `json:"group"`
	IAmTeacher        string   `json:"IAmTeacher"`
	IAmStudent        string   `json:"IAmStudent"`
	NextQuestion      string   `json:"nextQuestion"`
	PreviousQuestion  string   `json:"previousQuestion"`
	Login             string   `json:"login"`
	Password          string   `json:"password"`
	ShowPassword      string   `json:"showPassword"`
	RememberMe        string   `json:"rememberMe"`
	SignIn            string   `json:"signIn"`
	SignOut           string   `json:"signOut"`
	FirstForeignLang  jsonLang `json:"firstForeignLang"`
	SecondForeignLang jsonLang `json:"secondForeignLang"`
	ChooseSubject     string   `json:"chooseSubject"`
	ChooseTheme       string   `json:"chooseTheme"`
}
