package randvid

import (
    //"fmt"
    "net/http"
    "math/rand"
    "time"
)

func RandLink() string{
  VideoList := []string{
    "https://www.youtube.com/watch?v=tE1HDSipRxU",
    "http://img.4plebs.org/boards/f/image/1434/92/1434921182936.swf",
    "http://img.4plebs.org/boards/f/image/1442/53/1442530025364.swf",
    "http://img.4plebs.org/boards/f/image/1396/24/1396243585761.swf",
    "http://img.4plebs.org/boards/f/image/1421/12/1421126458648.swf",
    "https://www.youtube.com/watch?v=G2e_M06YDyY",
    "https://www.youtube.com/watch?v=DZGINaRUEkU",
    "https://www.youtube.com/watch?v=otnyM9RJG4o",
    "https://www.youtube.com/watch?v=2MN1vXO5JeI",
    "https://www.youtube.com/watch?v=u_jRgv-UqBU",
    "https://www.youtube.com/watch?v=DqC7H7_Noi8",
    "https://www.youtube.com/watch?v=2rEuie5lpGA",
    "http://i.4cdn.org/wsg/1444358103893.webm",
    "http://i.4cdn.org/wsg/1443246935183.webm",
    "http://i.4cdn.org/wsg/1443251102863.webm",
    "http://i.4cdn.org/wsg/1443254714327.webm",
    "http://i.4cdn.org/wsg/1443913579648.webm",
    "http://i.4cdn.org/wsg/1444112721405.webm",
    "http://i.4cdn.org/wsg/1443246656431.webm",
    "http://i.4cdn.org/wsg/1443199345935.webm",
    "http://i.4cdn.org/wsg/1443199345935.webm",
    "http://i.4cdn.org/wsg/1443199345935.webm",
    "http://i.4cdn.org/wsg/1443199345935.webm",
    "http://i.4cdn.org/wsg/1443199345935.webm",
    "http://i.4cdn.org/wsg/1443199345935.webm",
    "http://i.4cdn.org/wsg/1443199345935.webm",
    "http://i.4cdn.org/wsg/1443199345935.webm",
    "http://i.4cdn.org/wsg/1443199345935.webm",
    "http://i.4cdn.org/wsg/1443199345935.webm",
    "http://i.4cdn.org/wsg/1443199345935.webm",
    "http://i.4cdn.org/wsg/1443199345935.webm",
    "http://i.4cdn.org/wsg/1443199345935.webm",
    "http://i.4cdn.org/wsg/1443199345935.webm",
  }
  r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
  vid := r.Intn(len(VideoList)) + 1;
  return VideoList[vid]
}

func redirectHandler(path string) func(http.ResponseWriter, *http.Request) {
  return func (w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, path, http.StatusFound)
  }
}

func init() {
  http.HandleFunc("/", redirectHandler(RandLink()))
}

// goapp deploy -application rand-vid app.yaml
