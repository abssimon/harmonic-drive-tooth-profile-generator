/*!
 * Live Chat X, by Screets.
 *
 * SCREETS, d.o.o. Sarajevo. All rights reserved.
 * This  is  commercial  software,  only  users  who have purchased a valid
 * license  and  accept  to the terms of the  License Agreement can install
 * and use this program.
 */
'use strict';

(function() {
    var d = document;
    var b = d.createElement("script");
    b.src = lcxAPI.uri + '/assets/js/widget.js?v=' + Date.now();
    b.async = !0;
    d.head.appendChild(b);
})();