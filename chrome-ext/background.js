function resolved(record) {
  console.log(record.canonicalName);
  console.log(record.addresses);
}

var resolving = browser.dns.resolve(
"davidstclair.co.uk"
)
resolving.then(
  resolved
)

var tabId = 0
chrome.tabs.onActivated.addListener( function(activeInfo){
  chrome.tabs.get(activeInfo.tabId, function(tab){
    tabId = activeInfo.tabId;
    console.log("pony");
    console.log("you are here: " + activeInfo.tabId);
  });
});

var filter =  {urls: ["<all_urls>"]};
var opt_extraInfoSpec = [];

chrome.webRequest.onBeforeRequest.addListener(
  function(details) {
  console.log("Current: " + tabId)
  console.log("Tab Network Request: " + details.tabId)
  if ( tabId == details.tabId ) {
    console.log(details.url)
  }

  },
  filter,
  opt_extraInfoSpec
);