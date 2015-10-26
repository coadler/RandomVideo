package randvid

import (
    //"fmt"
    "net/http"
    "math/rand"
    "time"
    "html/template"

    "appengine"
    "appengine/datastore"
    "appengine/user"
)

type InsertLink struct {
  Video string
  Rand int
}

type Greeting struct {
        Author  string
        Content string
        Date    time.Time
}

var guestbookTemplate = template.Must(template.New("book").Parse(`
<html>
  <head>
    <title>Go Guestbook</title>
  </head>
  <body>
    {{range .}}
      {{with .Author}}
        <p><b>{{.}}</b> wrote:</p>
      {{else}}
        <p>An anonymous person wrote:</p>
      {{end}}
      <pre>{{.Content}}</pre>
    {{end}}
    <form action="/sign" method="post">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`))

const body = `
<html>
  <head>
    <title>Colin's Random-Video</title>
    <meta http-equiv="refresh" content="2; URL='{{.Video}}'" />
  </head>
  <body bgcolor="#ffffff">
    <center>
      Please wait to be redirected. If you are not redirected please click <a href="{{.Video}}"> here</a></br>
      Your random number is {{.Rand}}
    </center>
  </body>
</html>
`

func RandLink() (string, int){
  VideoList := []string{
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
    "http://data.desustorage.org/4ch/wsg/image/1443/69/1443692537150.webm", // 12 Thug Cat
    "http://data.desustorage.org/4ch/wsg/image/1443/24/1443246935183.webm", // 13 Plane Crash
    "http://data.desustorage.org/4ch/wsg/image/1443/25/1443251102863.webm", // 14 Old Spice
    "http://data.desustorage.org/4ch/wsg/image/1443/13/1443133300928.webm", // 15 McDonald's Remix
    "http://data.desustorage.org/4ch/wsg/image/1443/91/1443913579648.webm", // 16 Sledge Hammer
    "http://data.desustorage.org/4ch/wsg/image/1444/11/1444112721405.webm", // 17 Frozen Jesse Pinkman
    "http://data.desustorage.org/4ch/wsg/image/1441/07/1441076705814.webm", // 18 Talking Carl
    "http://data.desustorage.org/4ch/wsg/image/1443/10/1443109512298.webm", // 19 Mario Head Bang
    "http://data.desustorage.org/4ch/wsg/image/1443/03/1443034574003.webm", // 20 Gin+Juice
    "http://data.desustorage.org/4ch/wsg/image/1444/26/1444262136311.webm", // 21 Tomato -> Fan
    "http://data.desustorage.org/4ch/wsg/image/1444/26/1444262646541.webm", // 22 Old Man + Rollerskates
    "http://data.desustorage.org/4ch/wsg/image/1444/26/1444268978629.webm", // 23 Rugby Kid
    "http://data.desustorage.org/4ch/wsg/image/1440/94/1440943905455.webm", // 24 Beautiful Science
    "http://data.desustorage.org/4ch/wsg/image/1442/85/1442851024053.webm", // 25 Hitler Leek
    "http://data.desustorage.org/4ch/wsg/image/1444/31/1444314579514.webm", // 26 Bobcat Loading
    "http://data.desustorage.org/4ch/wsg/image/1444/33/1444339262566.webm", // 27 Throwing Knives
    "http://data.desustorage.org/4ch/wsg/image/1444/33/1444339554930.gif",  // 28 Perfect Circle
    "http://data.desustorage.org/4ch/wsg/image/1444/37/1444372368654.webm", // 29 Yodeling
    "http://data.desustorage.org/4ch/wsg/image/1444/41/1444415934593.webm", // 30 Leek Gun
    "http://data.desustorage.org/4ch/wsg/image/1444/22/1444228404320.webm", // 31 Cow Bell
    "http://data.desustorage.org/4ch/wsg/image/1444/46/1444467926220.webm", // 32 Arnold Palmer
    "https://www.youtube.com/watch?v=2svVkkNuSq0",                          // 33 Stop to my Beat
    "https://www.youtube.com/watch?v=DX_eeOZVS2o",                          // 34 Microwave Dance
    "http://data.desustorage.org/4ch/wsg/image/1444/17/1444175733369.webm", // 35 I Don't Need a Jacket
    "http://data.desustorage.org/4ch/wsg/image/1444/05/1444055325919.webm", // 36 GTFO
    "http://data.desustorage.org/4ch/wsg/image/1441/60/1441602608583.webm", // 37 Steven Universe
    "http://data.desustorage.org/4ch/wsg/image/1444/24/1444244530145.webm", // 38 BAD BOYZ
    "http://data.desustorage.org/4ch/wsg/image/1444/25/1444254081149.webm", // 39 BEER ME
    "http://data.desustorage.org/4ch/wsg/image/1444/37/1444371392998.webm", // 40 Pro Dad
    "http://data.desustorage.org/4ch/wsg/image/1442/75/1442751902230.webm", // 41 Drum Dog
    "http://data.desustorage.org/4ch/wsg/image/1440/91/1440917752642.webm", // 42 Get Down Cat
    "http://data.desustorage.org/4ch/wsg/image/1440/91/1440917996496.webm", // 43 Chicken Pokemon
    "http://data.desustorage.org/4ch/wsg/image/1443/81/1443810282577.webm", // 44 Dubstep Dog
    "http://data.desustorage.org/4ch/wsg/image/1442/15/1442157620314.webm", // 45 Shovel + Head
    "http://data.desustorage.org/4ch/wsg/image/1442/21/1442218215289.webm", // 46 Sly Kid
    "http://data.desustorage.org/4ch/wsg/image/1443/84/1443844978601.webm", // 47 Arriba
    "http://data.desustorage.org/4ch/wsg/image/1443/31/1443312573923.webm", // 48 Hot Boy Dog
    "http://i.4cdn.org/wsg/1440991924939.webm",                             // 49 More Doge
    "http://i.4cdn.org/wsg/1443393958007.webm",                             // 50 Even More Doge
    "http://i.imgur.com/J7VGU2g.gifv",                                      // 51 Kittens + Puppies
    "http://i.imgur.com/ZuMSuvM.gifv",                                      // 52 Doge on Ledge
    //"http://i.4cdn.org/wsg/1445033632154.webm",                           // 53 First Kiss/Life Insurance
    "http://i.4cdn.org/wsg/1445043480702.webm",                             // 54 Dancing Birdz
    "http://i.4cdn.org/wsg/1445058535377.webm",                             // 55 Wendy's Commercial
    "http://i.4cdn.org/wsg/1445068822578.webm",                             // 56 Water Bottle Kick
    "http://i.4cdn.org/wsg/1445110085377.webm",                             // 57 Pile of Balls
    "http://i.4cdn.org/wsg/1445110169319.webm",                             // 58 Mimicking Bird
    "http://i.4cdn.org/wsg/1445110618823.webm",                             // 59 Drum Keyboard
    "http://i.4cdn.org/wsg/1445110688973.webm",                             // 60 Rolaids
    "http://i.4cdn.org/wsg/1444956089069.webm",                             // 61 Terrible Email
    //"http://i.4cdn.org/wsg/1444966036876.webm",                           // 62 Trap + Horse
    "http://i.4cdn.org/wsg/1444985967218.webm",                             // 63 Racist SpongeBob
    "http://i.4cdn.org/wsg/1445031045338.webm",                             // 64 White Rapping
    "http://i.4cdn.org/wsg/1444877206633.webm",                             // 65 Amanda Berry Rap
    "http://i.4cdn.org/wsg/1444944916366.webm",                             // 66 Ukrainian Army Fail
    "http://i.4cdn.org/wsg/1444954291292.webm",                             // 67 Trump Dogg
    "http://i.4cdn.org/wsg/1444681180896.webm",                             // 68 Eminem Bill Cosby
    "http://i.4cdn.org/wsg/1444687230838.webm",                             // 69 Barbie March
    "http://i.4cdn.org/wsg/1444691894641.webm",                             // 70 Steal yo Girl
    "http://i.4cdn.org/wsg/1444697007201.webm",                             // 71 Jigsaw
    "http://i.4cdn.org/wsg/1444704274634.webm",                             // 72 British Jokes
    "http://i.4cdn.org/wsg/1444707961154.webm",                             // 73 Grinch Yoga
    "http://i.4cdn.org/wsg/1444719181433.webm",                             // 74 Chatty Patty
    "http://i.4cdn.org/wsg/1444728695622.webm",                             // 75 Lizard vs Cat
    "http://i.4cdn.org/wsg/1444729232374.webm",                             // 76 Raccoon
    //"",                             // 77
    //"",                             // 78
    //"",                             // 79
    //"",                             // 80
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

func guestbookKey(c appengine.Context) *datastore.Key {
        return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        q := datastore.NewQuery("Greeting").Ancestor(guestbookKey(c)).Order("-Date").Limit(10)
        greetings := make([]Greeting, 0, 10)
        if _, err := q.GetAll(c, &greetings); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        if err := guestbookTemplate.Execute(w, greetings); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
        }
}

func sign(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        g := Greeting {
                Content: r.FormValue("content"),
                Date:    time.Now(),
        }
        if u := user.Current(c); u != nil {
                g.Author = u.String()
        }
        key := datastore.NewIncompleteKey(c, "Greeting", guestbookKey(c))
        _, err := datastore.Put(c, key, &g)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        http.Redirect(w, r, "/suggest", http.StatusFound)
}

func init() {
  http.HandleFunc("/suggest", root)
  http.HandleFunc("/sign", sign)
  http.HandleFunc("/", Index)
}

// goapp deploy -application rand-vid app.yaml
