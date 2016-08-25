// Initialise system
var system = require('system');

// Initialise page
var page = require('webpage').create();

// Retrieve options
var url = system.args[1];

// Specify the encoding
phantom.outputEncoding = "UTF-8";

// Specify the user-agent
page.settings.userAgent = "Cirrius";
page.settings.loadImages = false;
page.settings.resourceTimeout = 3000;

// Open the specified url
page.open(url, function(status) {
    if (status === 'success') {
        page.evaluate(function() {
            Array.prototype.slice.call( document.getElementsByTagName("script") ).forEach(function(s) {
                s.parentNode.removeChild(s);
            });
        });
        console.log(page.content);
    }
    phantom.exit();
});
