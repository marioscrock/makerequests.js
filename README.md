# makerequests.js
A javascript library to generate requests on HTML elements with ```click()``` method defined. 

![Example UI](/img/makerequests.png)

## What makerequests.js can do? ##

* The interface let you choose the clickable object you want to make requests on or a ```random``` 
selection, for each request made, between all clickable elements specified.  

* The interface let you set a textual ```seed``` to initiate the random pseudo-generator

* If you do not select any distribution you can specify a ```Number of requests``` to be generated all at once.
If you select a distribution you have to provide parameters to determine how requests will be generated:
  * **Normal** (opt-name ```normal```)
    * ```Number of requests``` number of samples on the distribution
    * ```Mean``` and ```Standard Deviation``` to specify the distribution to be sampled
    * ```Time interval``` time interval in seconds to scale samples (min_sample = 0, max_sample = time_interval)
  * **Beta** (opt-name ```beta```)
    * ```Number of requests``` number of samples on the distribution
    * ```Alpha``` and ```Beta``` to specify the distribution to be sampled
    * ```Time interval``` time interval in seconds to scale samples (min_sample = 0, max_sample = time_interval)
  * **Chi Square** (opt-name ```chisquare```)
    * ```Number of requests``` number of samples on the distribution
    * ```Degrees of Freedom``` to specify the distribution to be sampled
    * ```Time interval``` time interval in seconds to scale samples (min_sample = 0, max_sample = time_interval)
  * **Exponential** (opt-name ```exp```)
    * ```Number of requests``` number of samples on the distribution
    * ```Rate``` to specify the distribution to be sampled
    * ```Time interval``` time interval in seconds to scale samples (min_sample = 0, max_sample = time_interval)
  * **Uniform** (opt-name ```uni```)
    * ```Number of requests``` number of samples on the distribution
    * ```Initial point a``` and ```Final point b``` to specify the distribution to be sampled
    * ```Time interval``` time interval in seconds to scale samples (min_sample = 0, max_sample = time_interval)
  * **T-Student** (opt-name ```studentT```)
    * ```Number of requests``` number of samples on the distribution
    * ```Degrees of Freedom``` to specify the distribution to be sampled
    * ```Time interval``` time interval in seconds to scale samples (min_sample = 0, max_sample = time_interval)
  * **Linear** (opt-name ```linear```)
    * ```Number of requests``` **Not** required
    * Each step (from 0 to ```Number of steps```) number of requests is increased linearly as specified by ```Slope``` parameter
    * ```Time interval``` time interval in seconds to scale requests (0, last_step = time_interval)
  * **Step Function** (opt-name ```step```)
    * ```Number of requests``` **Not** required
    * ```Low value``` and ```High value``` to specify number of requests on the high and low edge of the step
    * ```Ratio``` to specify the ratio between high and low edge of the step
    * ```Number of steps``` to specify the number of steps within a single repetition (high-low). It is rounded w.r.t
    the inserted ```Ratio```.
    * ```Repetitions``` to specify the number of high-low steps.
    * ```Time interval``` time interval in seconds to scale requests (0, last_step_of_last_repetition = time_interval)

**NOTE** Currently no checks are made on parameters inserted. All inputs are ```type='number'``` and empty 
ones results in a 0.  

* The library sets ```timeouts``` to generate requests as specified (and as sampled) and shows a graph reporting 
requests distribution over time.

## How to use makerequests.js ##
* Add class ```button-class-name``` to all HTML elements you need to generate requests on
* Add to HTML file an empty ```div``` element of class ```putMeHere-class-name container``` 

*Example*  
```xml
<body>
  <div class="putMeHere-class-name container"></div>
    <div class="col-md-3 col-sm-6"> <span class="btn btn-block button-class-name">Button 1</span></div>
    <div class="col-md-3 col-sm-6"> <span class="btn btn-block button-class-name">Button 2</span></div>
    <div class="col-md-3 col-sm-6"> <span class="btn btn-block button-class-name">Button 3</span></div>
    <div class="col-md-3 col-sm-6"> <span class="btn btn-block button-class-name">Button 4</span></div>
  </div>
</body>
```

* Add a tiny javascript snippet specifying:
  * ```buttonClassName``` class name identifying clickable elements
  * ```putMeHereClassName``` class name identifying the div container
  * The distributions to be enabled (*default*: no distribution shown)
  * Whether or not to visualize the graph of requests over time (*default*: graph not shown)
  * Whether or not to generate a .go file instead of calling ```click()``` on clickable elements (*default*: ```click()``` performed). See [section below](#generate-go-file) on how to use this option.

```javascript
var makeRequests = new MakeRequests({
  buttonClassName: "button-class-name",
  putMeHereClassName: "putMeHere-class-name",
  all: true,
  graph: true,
  generateGoFile: false});
  
makeRequests.build();
```

#### Select distributions ####

```javascript
all:true
``` 
The option above enables all distributions, otherwise you can add to parameters:
```javascript
opt-name:true
```
to selectively choose distributions to be enabled. ```opt-name``` of each distribution is specified above.

## Generate .go file ##

Modern browsers allow only a limited set of concurrent requests, so if your clickable elements performs an HTTP request when clicked, you can exploit the ```generateGoFile``` option to generate a ```.go``` file that if run together with the ```makeRequests.go``` file provided generates a set of go-routines to execute the requests as specified through the graphical interface.

To enable this option you need to:
* set ```generateGoFile: true``` while creating ```MakeRequests``` variable
* add to each clickable element an attribute ```url-key``` with value a unique string 
* set ```mapUrl: {"url-key-1":"url1", "url-key-2":"url-2, ..}``` while creating ```MakeRequests``` variable providing for each ```url-key``` the ```url``` to perform the HTTP request
	  
**Note** In case ```generateGoFile: true``` *random* option is not enabled.

When *Fire!* button is clicked a ```makeRequestsTimes.go``` is downloaded like the one here reported:

```go
package main

func main() {
	times := []float64{0,241.3264883321301, 9973.812807635615,10000}
	makeRequests(times,"http://example.url:8080/here?arg=value")
}
``` 

You can edit the ```makeRequests.go``` file provided to manage connections and set options for the HTTP request (e.g. headers). Then put both files in the same folder and run the go-routines with command:

```go run makeRequests.go makeRequestsTimes.go```

**Note** It requires golang ```v1.10```

### Requirements ###
* *bootstrap.js* version 4.0  
* *jquery.js* version 3.1.1
* *jstat.js* version 1.7.1
* *plotly.js* version 1.37.0
* *seedrandom.js* version 2.4.3

### Versioning ###
* ```v1.0``` First stable version of the library (*JSLint* compliant)
* ```v1.0.1``` Add opt for graph visualization
* ```v1.1``` Enable seed selection
* ```v1.2``` generateGoFile option
