# News API package for Go

A Golang client for the News API

##### Provided under MIT License by Daniel Vergara.
*Note: this library may be subtly broken or buggy. The code is released under
the MIT License – please take the following message to heart:*
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.


## Installation

```shell
go get github.com/dany2691/newsapigo
```

## Usage

After installation, import client into your project:

```go
import (
  "github.com/dany2691/newsapigo"
  )
```
## Initialization

```go
// Declare a map for the value of the params
params := make(map[string]string)
// Set a unique api key provided in https://newsapi.org/
apiKey := "XXXXXXXXXXXXXXXXXXXXXX"
var key = GetAPIKey()
key.SetKey(apiKey)
```
## Top Headlines

```go
  params["country"] = "us"
  newsapigo.GetTopHeadlines(params)
```

## Everything

```go
  params["q"] = "bitcoin"
  newsapigo.GetEverything(params)
```

## Sources

```go
  params["category"] = "general"
  newsapigo.GetSources(params)
```
