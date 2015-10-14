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
      Your random number is {{.Rand}}
    </center>
  </body>
</html>
`

const bodySFW = `
<html>
  <head>
    <title>Colin's Random-Video</title>
    <meta http-equiv="refresh" content="1; URL='{{.Video}}'" />
  </head>
  <body bgcolor="#ffffff">
    <center>
      Please wait to be redirected. If you are not redirected please click <a href="{{.Video}}"> here</a></br>
      New Videos are added every day. Check back often!</br>
      Your random number is {{.Rand}} </br>
      <b>YOU ARE USING THE SFW VERSION</b>
    </center>
  </body>
</html>
`

func RandLinkSFW() (string, int){
  VideoListSFW := []string{
    "https://www.youtube.com/watch?v=tE1HDSipRxU",                          // 0 Steve Irwin
    "http://img.4plebs.org/boards/f/image/1434/92/1434921182936.swf",       // 1  Just Do It
    "http://img.4plebs.org/boards/f/image/1442/53/1442530025364.swf",       // 2  SpongeBob
    //"http://img.4plebs.org/boards/f/image/1396/24/1396243585761.swf",       // 3  Big Shake
    //"http://img.4plebs.org/boards/f/image/1421/12/1421126458648.swf",       // 4  2012
    "https://www.youtube.com/watch?v=G2e_M06YDyY",                          // 5  Sieze the Day
    "https://www.youtube.com/watch?v=DZGINaRUEkU",                          // 6  Symphony of Science
    "https://www.youtube.com/watch?v=otnyM9RJG4o",                          // 7  Power of Music
    //"https://www.youtube.com/watch?v=2MN1vXO5JeI",                          // 8  Singalong
    "https://www.youtube.com/watch?v=u_jRgv-UqBU",                          // 9  Billy, Walmart Greeter
    "https://www.youtube.com/watch?v=DqC7H7_Noi8",                          // 10  Neil Armstrong Tribute
    "https://www.youtube.com/watch?v=2rEuie5lpGA",                          // 11 SURF THE NET
    "http://i.4cdn.org/wsg/1444358103893.webm",                             // 12 Thug Cat
    //"http://i.4cdn.org/wsg/1443246935183.webm",                             // 13 Plane Crash
    "http://i.4cdn.org/wsg/1443251102863.webm",                             // 14 Old Spice
    "http://i.4cdn.org/wsg/1443254714327.webm",                             // 15 McDonald's Remix
    "http://i.4cdn.org/wsg/1443913579648.webm",                             // 16 Sledge Hammer
    //"http://i.4cdn.org/wsg/1444112721405.webm",                             // 17 Frozen Jesse Pinkman
    "http://i.4cdn.org/wsg/1443246656431.webm",                             // 18 Talking Carl
    "http://i.4cdn.org/wsg/1443199345935.webm",                             // 19 Mario Head Bang
    //"http://i.4cdn.org/wsg/1443246507978.webm",                             // 20 Gin+Juice
    "http://i.4cdn.org/wsg/1444262136311.webm",                             // 21 Tomato -> Fan
    //"http://i.4cdn.org/wsg/1444262646541.webm",                             // 22 Old Man + Rollerskates
    "http://i.4cdn.org/wsg/1444268978629.webm",                             // 23 Rugby Kid
    "http://i.4cdn.org/wsg/1444276631906.webm",                             // 24 Beautiful Science
    //"http://i.4cdn.org/wsg/1444301298445.webm",                             // 25 Hitler Leek
    "http://i.4cdn.org/wsg/1444314579514.webm",                             // 26 Bobcat Loading
    "http://i.4cdn.org/wsg/1444339262566.webm",                             // 27 Throwing Knives
    "http://i.4cdn.org/wsg/1444339554930.gif",                              // 28 Perfect Circle
    "http://i.4cdn.org/wsg/1444372368654.webm",                             // 29 Yodeling
    //"http://i.4cdn.org/wsg/1444415934593.webm",                             // 30 Leek Gun
    "http://i.4cdn.org/wsg/1444448380633.webm",                             // 31 Cow Bell
    "http://i.4cdn.org/wsg/1444467926220.webm",                             // 32 Arnold Palmer
    //"https://www.youtube.com/watch?v=2svVkkNuSq0",                          // 33 Stop to my Beat
    "https://www.youtube.com/watch?v=DX_eeOZVS2o",                          // 34 Microwave Dance
    "http://i.4cdn.org/wsg/1443899505305.webm",                             // 35 I Don't Need a Jacket
    //"http://i.4cdn.org/wsg/1444055325919.webm",                             // 36 GTFO
    "http://i.4cdn.org/wsg/1441602608583.webm",                             // 37 Steven Universe
    "http://i.4cdn.org/wsg/1444244530145.webm",                             // 38 BAD BOYZ
    "http://i.4cdn.org/wsg/1444254081149.webm",                             // 39 BEER ME
    "http://i.4cdn.org/wsg/1444371392998.webm",                             // 40 Pro Dad
    "http://i.4cdn.org/wsg/1444440546084.webm",                             // 41 Drum Dog
    //"http://i.4cdn.org/wsg/1440917752642.webm",                             // 42 Get Down Cat
    "http://i.4cdn.org/wsg/1440917996496.webm",                             // 43 Chicken Pokemon
    "http://i.4cdn.org/wsg/1442157378169.webm",                             // 44 Dubstep Dog
    "http://i.4cdn.org/wsg/1442157620314.webm",                             // 45 Shovel + Head
    "http://i.4cdn.org/wsg/1442218215289.webm",                             // 46 Sly Kid
    "http://i.4cdn.org/wsg/1442384146512.webm",                             // 47 Arriba
    //"http://i.4cdn.org/wsg/1443312573923.webm",                             // 48 Hot Boy Dog
    "http://i.4cdn.org/wsg/1440991924939.webm",                             // 49 More Doge
    "http://i.4cdn.org/wsg/1443393958007.webm",                             // 50 Even More Doge
    //"",                                                                     // 51
  }
  r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
  vid := r.Intn(len(VideoListSFW))
  return VideoListSFW[vid], vid
}

func RandLinkNSFW() (string, int){
  VideoListNSFW := []string{
    "https://www.youtube.com/watch?v=tE1HDSipRxU",                          // 0 Steve Irwin
    "http://img.4plebs.org/boards/f/image/1434/92/1434921182936.swf",       // 1  Just Do It
    "http://img.4plebs.org/boards/f/image/1442/53/1442530025364.swf",       // 2  SpongeBob
    "http://img.4plebs.org/boards/f/image/1396/24/1396243585761.swf",       // 3  Big Shake
    "http://img.4plebs.org/boards/f/image/1421/12/1421126458648.swf",       // 4  2012
    "https://www.youtube.com/watch?v=G2e_M06YDyY",                          // 5  Sieze the Day
    "https://www.youtube.com/watch?v=DZGINaRUEkU",                          // 6  Symphony of Science
    "https://www.youtube.com/watch?v=otnyM9RJG4o",                          // 7  Power of Music
    "https://www.youtube.com/watch?v=2MN1vXO5JeI",                          // 8  Singalong
    "https://www.youtube.com/watch?v=u_jRgv-UqBU",                          // 9  Billy, Walmart Greeter
    "https://www.youtube.com/watch?v=DqC7H7_Noi8",                          // 10  Neil Armstrong Tribute
    "https://www.youtube.com/watch?v=2rEuie5lpGA",                          // 11 SURF THE NET
    "http://i.4cdn.org/wsg/1444358103893.webm",                             // 12 Thug Cat
    "http://i.4cdn.org/wsg/1443246935183.webm",                             // 13 Plane Crash
    "http://i.4cdn.org/wsg/1443251102863.webm",                             // 14 Old Spice
    "http://i.4cdn.org/wsg/1443254714327.webm",                             // 15 McDonald's Remix
    "http://i.4cdn.org/wsg/1443913579648.webm",                             // 16 Sledge Hammer
    "http://i.4cdn.org/wsg/1444112721405.webm",                             // 17 Frozen Jesse Pinkman
    "http://i.4cdn.org/wsg/1443246656431.webm",                             // 18 Talking Carl
    "http://i.4cdn.org/wsg/1443199345935.webm",                             // 19 Mario Head Bang
    "http://i.4cdn.org/wsg/1443246507978.webm",                             // 20 Gin+Juice
    "http://i.4cdn.org/wsg/1444262136311.webm",                             // 21 Tomato -> Fan
    "http://i.4cdn.org/wsg/1444262646541.webm",                             // 22 Old Man + Rollerskates
    "http://i.4cdn.org/wsg/1444268978629.webm",                             // 23 Rugby Kid
    "http://i.4cdn.org/wsg/1444276631906.webm",                             // 24 Beautiful Science
    "http://i.4cdn.org/wsg/1444301298445.webm",                             // 25 Hitler Leek
    "http://i.4cdn.org/wsg/1444314579514.webm",                             // 26 Bobcat Loading
    "http://i.4cdn.org/wsg/1444339262566.webm",                             // 27 Throwing Knives
    "http://i.4cdn.org/wsg/1444339554930.gif",                              // 28 Perfect Circle
    "http://i.4cdn.org/wsg/1444372368654.webm",                             // 29 Yodeling
    "http://i.4cdn.org/wsg/1444415934593.webm",                             // 30 Leek Gun
    "http://i.4cdn.org/wsg/1444448380633.webm",                             // 31 Cow Bell
    "http://i.4cdn.org/wsg/1444467926220.webm",                             // 32 Arnold Palmer
    "https://www.youtube.com/watch?v=2svVkkNuSq0",                          // 33 Stop to my Beat
    "https://www.youtube.com/watch?v=DX_eeOZVS2o",                          // 34 Microwave Dance
    "http://i.4cdn.org/wsg/1443899505305.webm",                             // 35 I Don't Need a Jacket
    "http://i.4cdn.org/wsg/1444055325919.webm",                             // 36 GTFO
    "http://i.4cdn.org/wsg/1441602608583.webm",                             // 37 Steven Universe
    "http://i.4cdn.org/wsg/1444244530145.webm",                             // 38 BAD BOYZ
    "http://i.4cdn.org/wsg/1444254081149.webm",                             // 39 BEER ME
    "http://i.4cdn.org/wsg/1444371392998.webm",                             // 40 Pro Dad
    "http://i.4cdn.org/wsg/1444440546084.webm",                             // 41 Drum Dog
    "http://i.4cdn.org/wsg/1440917752642.webm",                             // 42 Get Down Cat
    "http://i.4cdn.org/wsg/1440917996496.webm",                             // 43 Chicken Pokemon
    "http://i.4cdn.org/wsg/1442157378169.webm",                             // 44 Dubstep Dog
    "http://i.4cdn.org/wsg/1442157620314.webm",                             // 45 Shovel + Head
    "http://i.4cdn.org/wsg/1442218215289.webm",                             // 46 Sly Kid
    "http://i.4cdn.org/wsg/1442384146512.webm",                             // 47 Arriba
    "http://i.4cdn.org/wsg/1443312573923.webm",                             // 48 Hot Boy Dog
    "http://i.4cdn.org/wsg/1440991924939.webm",                             // 49 More Doge
    "http://i.4cdn.org/wsg/1443393958007.webm",                             // 50 Even More Doge
    //"",                                                                     // 51
  }
  r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
  vid := r.Intn(len(VideoListNSFW))
  return VideoListNSFW[vid], vid
}

func IndexSFW(w http.ResponseWriter, r *http.Request) {
  w.Header().Add("Content Type", "text/html")
  tmpl, err := template.New("video").Parse(bodySFW)
  video, rand := RandLinkSFW()
  if err == nil {
    redirect := InsertLink{video, rand}
    tmpl.Execute(w, redirect)
  } else {
    panic(err)
  }
}

func IndexNSFW(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("video").Parse(body)
    video, rand := RandLinkNSFW()
    if err == nil {
      redirect := InsertLink{video, rand}
      tmpl.Execute(w, redirect)
    } else {
      panic(err)
    }
}

func init() {
  http.HandleFunc("/sfw", IndexSFW)
  http.HandleFunc("/.*", IndexNSFW)
}

// goapp deploy -application rand-vid app.yaml
