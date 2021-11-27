/**
 * Live Chat X, by Screets.
 *
 * SCREETS, d.o.o. Sarajevo. All rights reserved.
 * This  is  commercial  software,  only  users  who have purchased a valid
 * license  and  accept  to the terms of the  License Agreement can install
 * and use this program.
 */

'use strict';

var lcx_events, lcx_frontend;

(function () {

  var xhr = new XMLHttpRequest();
  var fd = new FormData();
  var data = {
    mode: 'getWidget',
    action: 'lcx_action',
    _ajax_nonce: lcxAJAX.nonce
  };

  xhr.open('POST', lcxAJAX.uri, true);

  // Handle response
  xhr.onreadystatechange = function () {

    if (xhr.readyState == 4) {

      // Perfect!
      if (xhr.status == 200) {
        if (cb) {
          cb(JSON.parse(xhr.responseText));
        }

        // Something wrong!
      } else {
        if (cb) {
          cb(null);
        }
      }
    }
  };

  // Get data
  for (var k in data) {
    fd.append(k, data[k]);
  }var cb = function cb(r) {
    // Load assets
    if (r.assets) {
      var obj;
      var total = r.assets.length;
      var head = document.getElementsByTagName('head')[0];
      for (var i = 0; i < total; i++) {

        // JS
        if (r.assets[i].type === 'js') {
          obj = document.createElement('script');
          obj.async = false;
          obj.src = r.assets[i].src;
          document.head.appendChild(obj);

          // CSS
        } else {
          obj = document.createElement('link');
          obj.rel = 'stylesheet';
          obj.type = 'text/css';
          obj.href = r.assets[i].href;
          head.appendChild(obj);
        }

        if (i + 1 === total) {
          // last asset
          obj.addEventListener('load', function () {
            init(r);
          });
        }
      }
    }
  };

  var init = function init(r) {

    var classes = 'lcx lcx-widget';

    if (r.opts.hideOffline) classes += ' lcx--hidden';

    // Create widget object
    var widget = document.createElement('div');
    widget.id = 'lcx-widget';
    widget.className = classes;
    document.body.appendChild(widget);

    // Create an iframe
    var iframe = document.createElement('iframe');
    iframe.setAttribute('allowfullscreen', '');
    widget.appendChild(iframe);

    // Write iframe content
    var ibody = iframe.contentWindow.document;
    ibody.open();
    ibody.write(r.iframe);
    ibody.close();

    var opts = {
      _iframe: iframe, //include iframe object
      db: {},
      ajax: {},
      user: {},
      autoinit: true,
      initPopup: 'online',
      ntfDuration: 5000, // ms.
      platform: 'frontend',
      dateFormat: 'd/m/Y',
      hourFormat: 'H:i',

      // Company data
      companyName: '',
      companyURL: '',
      companyLogo: '',
      anonymousImage: '',
      systemImage: ''
    };

    for (var k in r.opts) {
      opts[k] = r.opts[k];
    }lcx_events = new nBirdEvents();
    lcx_frontend = new nightBird(opts, r.strings);

    // Load extra assets if exists
    if (r.extraAssets) {
      var obj = void 0;
      var _iteratorNormalCompletion = true;
      var _didIteratorError = false;
      var _iteratorError = undefined;

      try {
        for (var _iterator = r.extraAssets[Symbol.iterator](), _step; !(_iteratorNormalCompletion = (_step = _iterator.next()).done); _iteratorNormalCompletion = true) {
          var file = _step.value;

          // JS
          if (file.type === 'js') {
            obj = document.createElement('script');
            obj.async = false;
            obj.src = file.src;
            document.head.appendChild(obj);

            // CSS
          } else {
            obj = document.createElement('link');
            obj.rel = 'stylesheet';
            obj.type = 'text/css';
            obj.href = file.href;
            head.appendChild(obj);
          }
        }
      } catch (err) {
        _didIteratorError = true;
        _iteratorError = err;
      } finally {
        try {
          if (!_iteratorNormalCompletion && _iterator.return) {
            _iterator.return();
          }
        } finally {
          if (_didIteratorError) {
            throw _iteratorError;
          }
        }
      }
    }
  };

  // Initiate a multipart/form-data upload
  xhr.send(fd);
})();