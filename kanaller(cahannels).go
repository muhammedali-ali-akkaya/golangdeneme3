package main

func main() {
	k := make(chan bool) //kannalar make fonkisiyonu ile oluşturulur
	k <- true            //kanala veri atama işlemi
	a := <-k             //kanaldaki degişkeni degere atadık
	<-k                  //sadece kanaladeger gelmesini beklemek

}
