Vegeta is a http load testing tool implemented with Go. If you haven't looked at it yet here is the github link [Vegeta](https://github.com/tsenart/vegeta).

I liked the code and learned some cool techniques, so here I share.

In Vegeta there is a structure called `Target` which represents a request endpoint as you can tell from the type definition:

<pre class="prettyprint lang-go">
// Just holds request information 
type Target struct {
	Method string
	URL    string
	Body   []byte
	Header http.Header
}

// http.Request object is generated eventually
func (t *Target) Request() (*http.Request, error) {...}
</pre>

Typically the Vegeta attacker (who sends the http requests) will take a list of Targets and create `http.Request` and send out.

What is interesting is that a `Target` can be generated eagerly and lazily. To achieve this a function generator pattern is used.

Think of generator as an iterator but instead of calling `x.next()` to get next item, the generator just calls itself 
thanks to the fact that in Go function is a pretty flexible that can be typed, assigned and also be context-aware (closure).

Let me make this more clear by pasting some code here:

<pre class="prettyprint lang-go">
// Targetor is a generator function
type Targeter func() (*Target, error)

func NewEagerTargeter(...) (Targeter, error) {...}
func NewLazyTargeter(...) (Targeter, error) {...}
</pre>

As you can tell, `NewEagerTargeter`, `NewLazyTargeter` are used as constructor for generator functions and
 also note that the generator is typed for a better readability. The constructor could return a untyped generator function but less readable instead like following:

<pre class="prettyprint lang-go">
func NewLazyTargeter(...) (func() (*Target, error), error) {...}
</pre>

So now, how does the generator its self been constructed? It becomes very clear when look at the implementation of the lazy targeter e.g.

<pre class="prettyprint lang-go">
func NewLazyTargeter(src io.Reader, body[], hdr http.Header) (Targeter, error) {
  ...
  sc := peekingScanner{src: bufio.NewScanner(src)}

  return function() (*Target, error) {
    ...
    sc.Scan()    
    ...
  }
}
</pre>

The code uses [golang's function closure](https://tour.golang.org/moretypes/21) to be stateful about what 
the current scanning state is (thus track which line we are at right now).

The rest of the code inside the closure function is just to parse the scanned lines and determine if more lines need to be scanned to generate one `Target` object.

Pretty slick!

_Stay tuned for more Golang stuff!_
