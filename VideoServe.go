package randvid

import (
    //"fmt"
    "net/http"
    "math/rand"
    "time"
    "html/template"
)

type InsertLink struct {
  Video string
  Rand int
}

const body = `
<html>
  <head>
    <title>Colin's Random-Video</title>
    <meta http-equiv="refresh" content="1; URL='{{.Video}}'" />
  </head>
  <body bgcolor="#ffffff">
    <center>
      Please wait to be redirected. If you are not redirected please click <a href="{{.Video}}"> here</a></br>
      New Videos are added every day. Check back often!</br>
      Investigating problems with how random the videos actually are. Results seem extremely skewed.</br>
      Your random number is {{.Rand}}
    </center>
  </body>
</html>
`

func RandLink() (string, int){
  VideoList := []string{
    "https://www.youtube.com/watch?v=tE1HDSipRxU",
    "http://img.4plebs.org/boards/f/image/1434/92/1434921182936.swf",       // 0  Steve Irwin
    "http://img.4plebs.org/boards/f/image/1442/53/1442530025364.swf",       // 1  Just Do It
    "http://img.4plebs.org/boards/f/image/1396/24/1396243585761.swf",       // 2  SpongeBob
    "http://img.4plebs.org/boards/f/image/1421/12/1421126458648.swf",       // 3  Big Shake
    "https://www.youtube.com/watch?v=G2e_M06YDyY",                          // 4  2012
    "https://www.youtube.com/watch?v=DZGINaRUEkU",                          // 5  Sieze the day
    "https://www.youtube.com/watch?v=otnyM9RJG4o",                          // 6  Symphony of Science
    "https://www.youtube.com/watch?v=2MN1vXO5JeI",                          // 7  Power of Music
    "https://www.youtube.com/watch?v=u_jRgv-UqBU",                          // 8  Singalong
    "https://www.youtube.com/watch?v=DqC7H7_Noi8",                          // 9  Billy, Walmart Greeter
    "https://www.youtube.com/watch?v=2rEuie5lpGA",                          // 10 Neil Armstrong Tribute
    "http://i.4cdn.org/wsg/1444358103893.webm",                             // 11 SURF THE NET
    "http://i.4cdn.org/wsg/1443246935183.webm",                             // 12 Thug Cat
    "http://i.4cdn.org/wsg/1443251102863.webm",                             // 13 Plane Crash
    "http://i.4cdn.org/wsg/1443254714327.webm",                             // 14 Old Spice
    "http://i.4cdn.org/wsg/1443254714327.webm",                             // 15 McDonald's Remix
    "http://i.4cdn.org/wsg/1443913579648.webm",                             // 16 Sledge Hammer
    "http://i.4cdn.org/wsg/1444112721405.webm",                             // 17 Frozen Jesse Pinkman
    "http://i.4cdn.org/wsg/1443246656431.webm",                             // 18 Talking Carl
    "http://i.4cdn.org/wsg/1443199345935.webm",                             // 19 Mario Head Bang
    "http://i.4cdn.org/wsg/1443246507978.webm",                             // 20 Gin+Juice
    "http://i.4cdn.org/wsg/1444262136311.webm",                             // 21 Tomato -> Fan
    "http://i.4cdn.org/wsg/1444262646541.webm",                             // 22 Old Man + Rollerskates
    "http://i.4cdn.org/wsg/1444268978629.webm",                             // 23 Rugby Kid
    "http://i.4cdn.org/wsg/1444276631906.webm",                             // 25 Beautiful Science
    "http://i.4cdn.org/wsg/1444301298445.webm",                             // 25 Hitler Leek
    "http://i.4cdn.org/wsg/1444314579514.webm",                             // 26 Bobcat Loading
    "http://i.4cdn.org/wsg/1444339262566.webm",                             // 27 Throwing Knives
    "http://i.4cdn.org/wsg/1444339554930.gif",                              // 28 Perfect Circle
    "http://i.4cdn.org/wsg/1444372368654.webm",                             // 29 Yodeling
    "http://i.4cdn.org/wsg/1444415934593.webm",                             // 30 Leek Gun
    "http://i.4cdn.org/wsg/1444448380633.webm",                             // 31 Cow Bell
    "http://i.4cdn.org/wsg/1444467926220.webm",                             // 32 Arnold Palmer
  }
  r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
  vid := r.Intn(len(VideoList))
  return VideoList[vid], vid
}

func Index(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("video").Parse(body)
    video, rand := RandLink()
    if err == nil {
      redirect := InsertLink{video, rand}
      tmpl.Execute(w, redirect)
    } else {
      panic(err)
    }
}

func init() {
  http.HandleFunc("/", Index)
}

// goapp deploy -application rand-vid app.yaml
