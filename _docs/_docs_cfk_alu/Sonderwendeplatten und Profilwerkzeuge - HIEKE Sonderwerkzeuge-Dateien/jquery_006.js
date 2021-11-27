/* strech jQuery Plugin
 *
 * Copyright (c) 2010, Tom Switzer <thomas.switzer@gmail.com>
 *
 * Permission to use, copy, modify, and/or distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */
(function($) {

/**
 * Expands text to fill up the entire width of its parent (or at least as much
 * as possible). It does this in 2 ways. First, it finds the largest font-size
 * for the text that can still fits on one line. If there is still some space
 * left to fill, it will increase the gap between the words as well to fill the
 * remaining space.
 *
 * For example, say we wanted to make the following title, "Sprockets and 
 * Widgets" take up as much of the full 400px of width as possible:
 *
 *  <h1 style="width: 400px">Sprockets and Widgets</h1>
 *  <script>
 *    $("h1").contents().fillWidth();
 *  </script>
 */
$.fn.stretch = function(opts) {
	opts = $.extend({}, $.fn.stretch.defaults, opts);
	if (!(opts.max >= 0) || !(opts.min >= 0))
		opts.min = opts.max = 0;

    this.each(function() {
        var container = $(this),
            contents = container.find("> .stretch--handle");
        
        if (!container.hasClass("stretch--resizer") || !contents) {
            contents = $(this).wrap("<span class='stretch--handle' />").parent();
            container = contents.wrap("<div class='stretch--resizer' />").parent();
            contents.css("margin", "0").css("padding", "0");
            container.css("margin", "0").css("padding", "0");
            
            // Ensures we only work with 1 line.
        
            container.css("white-space", "nowrap").css("overflow", "hidden");
        }
        
        var idealWidth = container.width(),
            width,
            min = opts.min || 1,
            max = min;
        
        // Search for a minimum/maximum font size so we can bound our search.
        
		if (!opts.max) {
			do {
				min = max;
				max *= 2;
				container.css("font-size", max + "px");
				
				// If the width isn't changing, then avoid an infinite loop.
				var realWidth = contents.width();
				width = realWidth <= width ? idealWidth : realWidth;
				
			} while (width < idealWidth);
		} else {
			max = opts.max;
			if (min == max)
				container.css("font-size", max + "px");
		}
        
        if (width == idealWidth)
            return;
        
        // Perform a binary search to find the font-size closest to the ideal.
        
        while (min < max) {
            var c = Math.floor((min + max) / 2);
            container.css("font-size", c + "px");
            width = contents.width();
            if (width == idealWidth)
                break;
            if (width < idealWidth)
                min = c + 1;
            else
                max = c;
        }
        
        // We may end up on the wrong side of "perfect," so we fix this.
        
        if (width > idealWidth)
            container.css("font-size", (max - 1) + "px");
        
        // Adding a little bit of word spacing will bring us that much closer.
        
        var spacing = 0,
            origWidth = contents.width(),
			maxSpacing = opts.maxSpacing;
        do {
            spacing += 1;
            container.css("word-spacing", spacing + "px");
            width = contents.width();
        } while (spacing <= maxSpacing && width <= idealWidth && width > origWidth);
        container.css("word-spacing", (spacing - 1) + "px");
    });
    
    return this;
};

$.fn.stretch.defaults = {
	min: 0,
	max: 0,
	maxSpacing: 0
};
    
})(jQuery);
