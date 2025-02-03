package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var AnsArray = Answers{
	Ans: []Answer{{
		First:  0,
		Second: 0,
		Third:  0,
		Fourth: 1,
	}},
}

func Game(w http.ResponseWriter, r *http.Request) {

	n1, _ := strconv.Atoi(r.FormValue("first"))
	n2, _ := strconv.Atoi(r.FormValue("second"))
	n3, _ := strconv.Atoi(r.FormValue("third"))
	n4, _ := strconv.Atoi(r.FormValue("fourth"))

	
	Ans := Answer{
		First:  n1,
		Second: n2,
		Third:  n3,
		Fourth: n4,
		//add markers
	}

	if !ValidateAnswer(Ans) {
		t, err := template.New("loose.html").ParseFiles("public/templates/loose.html")
		if err != nil {
			log.Fatal(err)
		}
		err = t.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	if Ans.First == 0 && Ans.Second == 0 && Ans.Third == 0 && Ans.Fourth == 0 {

		t, err := template.New("win.html").ParseFiles("public/templates/win.html")
		if err != nil {
			log.Fatal(err)
		}
		err = t.Execute(w, AnsArray)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		//store the values and continue
		AnsArray.Ans = append(AnsArray.Ans, Ans)
		t, err := template.New("index.html").ParseFiles("public/templates/index.html")
		if err != nil {
			log.Fatal(err)
		}
		err = t.Execute(w, Ans)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func ValidateAnswer(A Answer) bool {
	
	s1 := strconv.Itoa(A.First)
	s2 := strconv.Itoa(A.Second)
	s3 := strconv.Itoa(A.Third)
	s4 := strconv.Itoa(A.Fourth)

	if len(s1) != 1 || len(s2) != 1 || len(s3) != 1 || len(s4) != 1 {
		return false
	}

	if (A.First != 0 && A.First != 1 && A.First != 2) || (A.Second != 0 && A.Second != 1 && A.Second != 2) || (A.Third != 0 && A.Third != 1 && A.Third != 2) || (A.Fourth != 0 && A.Fourth != 1 && A.Fourth != 2) {
		return false
	}
	return true
}
