// MAIN JS FILE
// VERSION:		 	2.0
// LAST UPDATE: 	17.03.17



// SET STRICT MODE //
"use strict";



// VARS //
var linkLocation = "#";
var isMobile = false;
var redirectTime = 350;
var doNotUnload = false;



// EVENTS //
$(document).ready(function(){
	// Temp Fix
	//$(".mod_article").addClass("span_12_of_12");
	//$(".mod_article").addClass("first");

	// Set Lang nav DOM position
	$(".mod_changelanguage").prependTo("#languageSwitch");
	$(".mod_changelanguage").clone().prependTo("#languageSwitchHeader");

	// LOAD FUNCTIONS
	smoothLoad();
	breadCrumb();
	nav();
	wrapLinks();
	hoverEffects();
	//contactForm();
	//redirectAnimations();
	videos();
	newsList();
	slideShow();


	// IF MAP EXISTS LOAD FUNCTION
	if ($("#map").length) {
		googleMap();
	}

	$(window).resize();
	$(window).scroll();

	// CHECK IF MOBILE
	if( /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent) ) {
		isMobile = true;
		$("body").addClass("mobile");
	} else {
		isMobile = false;
		$("body").addClass("desktop");
	}

	// SET HEIGHTS
	setHeight();




	// Forms
	$(".datePicker input").datepicker({
		language: 'de-DE',
		format: 'dd.mm.yyyy',
		autoHide: true
	});


	$(".onlyOneCheckbox").each(function(){
		var realThis = $(this);

		realThis.find("input").on('change', function () {
			realThis.find("input").not(this).prop("checked", false);
		});
	});
	
	
	// Form
	$(".fraeser label").click(function(e) {
		// e.preventDefault();
		var fraeser = $(this).parents(".fraeser");
		// console.log(fraeser.attr("class"));
		
		if (fraeser.hasClass("active")) {
			fraeser.removeClass("active");
			$(".fraeser").show(250);
		} else {
			fraeser.addClass("active");
			$(".fraeser").not(".active").hide(250);
			fraeser.show(250);
		}
		
	});
	
	
	// Form
	$(".schaft label").click(function(e) {
		// e.preventDefault();
		var fraeser = $(this).parents(".schaft");
		// console.log(fraeser.attr("class"));
		
		if (fraeser.hasClass("active")) {
			fraeser.removeClass("active");
			$(".schaft").show(250);
		} else {
			fraeser.addClass("active");
			$(".schaft").not(".active").hide(250);
			fraeser.show(250);
		}
		
	});
	


});


// BACK/FORWARD EVENT
window.onpageshow = function(event) {
    if (event.persisted) {
        //smoothLoad();
    }
};


// UNLOAD EVENT
/*
window.onbeforeunload = function(){
	if (doNotUnload == false) {
		$("body").addClass("leavePage");
		$("#navIcon, #mainNav, header").removeClass("active");
		$("#main").removeClass("fadeIn").addClass("fadeOut");
		$("footer").fadeTo(200, 0);
	}
};
*/


// WINDOW LOAD EVENT
$(window).on('load', function () {
	fixes();
	setHeight();

	$(window).resize();
	window.dispatchEvent(new Event('resize'));
});


// SCROLL EVENT
$(window).scroll(function() {

	// FADE CONTENT IN ON SHOW
	$('.showOnScroll').each( function(i){
		var divBottom = $(this).offset().top + $(this).outerHeight() * 0.4;
		var windowBottom = $(window).scrollTop() + $(window).height();

		if( windowBottom > divBottom ){
			$(this).animate({'opacity':'1'}, 450);   
		}
	});
	

});

$(".tableWrapper").scroll(function(){
	$(this).addClass("scrolled");
});


// RESIZE EVENT
$(window).resize(function(){
	if ($(window).width() < 800) {
		redirectTime = 550;
	} else {
		redirectTime = 350;
	}

	setHeight();
	
	resetHoverEffects();
	fixes();

	$(".roundBox").each(function(){
		$(this).css("height", $(this).width());
		$(this).find("h2").css("margin", "auto");
		$(this).find("h2").css("padding-top", ($(this).height() - $(this).find("h2").height()) / 2 + "px");
	});


	$(".ce_previewdownload").each(function(){
		var nextHeight = $(this).next(".ce_previewdownload").outerHeight();

		if ($(this).height() < nextHeight) {
			$(this).css("height", nextHeight);
		} else {
			$(this).next(".ce_previewdownload").css("height", $(this).height());
		}

	});
});


// CLICK EVENT
$(document).click(function(event){
	hideMenuOverlays(event);
});


// KEY DOWN EVENT
$(document).keydown(function(e) {
     if (e.keyCode == 27) {
        hideNav();
    }
});



// FUNCTIONS //

// SMOOTH LOADING FUNCTION
function smoothLoad() {
	// FIX FOOTER
	$("footer").insertAfter("#wrapper");

	// FADE CONTENT IN
	/*
	setTimeout(function(){
		$("#main").addClass("fadeIn");
	}, 50);

	setTimeout(function(){
		$("footer").fadeTo(300, 1).css("display", "inline-block");
		$("#main").addClass("normal");
		$(window).resize();
	}, 400);
	*/

	$(window).resize();
}


// NAV
function nav() {
	// OPEN MAIN NAV ON BURGER CLICK
	$("#navBurger").click(function() {
		toggleNav();
	});

	// SET BURGER HOVER CLASS
	$("#navBurger").mouseenter(function(){
		$("#navBurger, #navIcon").addClass("hover");
	});

	$("#navBurger").mouseleave(function(){
		$("#navBurger, #navIcon").removeClass("hover");
	});


	// Nav white background
	var $document = $(document),
		$element = $("#header"),
		className = "sticky";

	$document.scroll(function () {
		$element.toggleClass(className, $document.scrollTop() >= 30);
	});
}


// HIDE NAV
function hideNav() {
	$("#mainNav, header, #mainNav .inner, #main").removeClass("active");

	// REMOVE NAV AND BODY ACTIVE CLASS
	$("body, #navIcon").removeClass("active");
}


// SHOW NAV
function showNav() {
	$("#mainNav, header, #navIcon, #mainNav .inner, #main, body").addClass("active");
}


// TOGGLE NAV
function toggleNav() {
	// TOGGLE HEADER
	$("#main").removeClass("fadeIn");
	$("#mainNav, header, #main, body").toggleClass("active");

	// TOGGLE NAV CONTENT
	setTimeout(function(){
		$("#mainNav .inner").toggleClass("active");
	}, 200);

	// TOGLE NAV ICON
	$("#navIcon").toggleClass("active");
}


// REDIRECT ANIMATIONS
function redirectAnimations() {
    $(document).on( "click", "a", function(event) {
    	if (!$(this).hasClass("slider-prev") && !$(this).hasClass("slider-next")) {

	        if ($(this).attr("target") !== "_blank" && !$(this).hasClass("email") ) {
	        	event.preventDefault();
	        	doNotUnload = true;

	        	linkLocation = this.href;
	        	$("*").css("cursor", "wait");
	    		$("#navIcon, #mainNav, header").removeClass("active");
	        	$("#main").removeClass("fadeIn").addClass("fadeOut");
	        	$("footer").fadeTo(300, 0);
	        	
	        	setTimeout(function(){
					redirectPage();
				}, redirectTime);
	        } else {
	        	event.preventDefault();
	        	doNotUnload = true;
	        	linkLocation = this.href;
	        	window.open(linkLocation, "_blank");

	        	setTimeout(function(){
					doNotUnload = false;
				}, 150);
	        }

	    }
    });

    function redirectPage() {
        window.location = linkLocation;
    }
}


// HIDE OVERLAYS ON CLICK
function hideMenuOverlays(event) {
    var target = event.target;

    if (!$(target).is('#mainNav') && !$(target).is('#navBurger') && !$(target).parents().is('#mainNav') && !$(target).parents().is('#navBurger')) {
        hideNav();
    }
}


// WRAP DIV IN LINKS
function wrapLinks() {
	$(".overlayBox, .roundBox").each(function(){
		var link = $(this).find(".overlayLink").attr("href");
		var classes = $(this).find(".overlayLink").attr("class");

		if (link !== undefined) {
			$(this).wrap('<a href="' + link + '" class="boxLink ' + classes + '"></a>"');
		}
	});


	$(".fullWidthBox").each(function () {
		var link = $(this).find(".arrowLink a").attr("href");
		var classes = $(this).find(".arrowLink a").attr("class");

		if (link !== undefined) {
			$(this).wrap('<a href="' + link + '" class="boxLink ' + classes + '"></a>"');
		}
	});


	$(".slideImg p").each(function(){
		var link = $(this).parent().find("h3 a").attr("href");
		var classes = $(this).parent().find("h3 a").attr("class");

		if (link !== undefined) {
			$(this).wrap('<a href="' + link + '" class="boxLink ' + classes + '"></a>"');
		}
	});

	$(".mod_nl_list li").each(function(){
		var link = $(this).find("a").attr("href");

		if (link !== undefined) {
			$(this).wrap('<a href="' + link + '" class="boxLink newsletterLink"></a>"');
		}
	});
}


// HOVER EFFECTS
function hoverEffects() {
	// ON MOUSE ENTER
	$(".boxLink").click(function(){
		// IMG STYLES
		var img = $(this).find(".bgPic img");
		var h2 = $(this).find("h2, .headline");

		$(img).addClass("animate");

		setTimeout(function(){
			$(h2).addClass("hover");
		}, 300);

	});


	$(".slideBox").click(function(){
		// IMG STYLES
		var img = $(this).find(".bgPic img");
		var h2 = $(this).find("h3");

		$(img).addClass("animate");

		setTimeout(function(){
			$(h2).addClass("hover");
		}, 300);

	});


	$(".boxLink").mouseenter(function(){
		// IMG STYLES
		var img = $(this).find(".bgPic img");
		var h2 = $(this).find("h2, .headline");

		$(img).addClass("animate");

		setTimeout(function(){
			$(h2).addClass("hover");
		}, 300);

	});

	$(".slideBox").mouseenter(function(){
		// IMG STYLES
		var img = $(this).find(".bgPic img");
		var h2 = $(this).find("h3");

		$(img).addClass("animate");

		setTimeout(function(){
			$(h2).addClass("hover");
		}, 300);

	});

	// ON MOUSE LEAVE
	$(".boxLink").mouseleave(function(){
		// IMG STYLES
		var img = $(this).find(".bgPic img");
		var h2 = $(this).find("h2, .headline");

		$(img).removeClass("animate");
		$(h2).removeClass("hover");

		setTimeout(function(){
			$(h2).removeClass("hover");
		}, 300);
	});

	$(".slideBox").mouseleave(function(){
		// IMG STYLES
		var img = $(this).find(".bgPic img");
		var h2 = $(this).find("h3");

		$(img).removeClass("animate");
		$(h2).removeClass("hover");

		setTimeout(function(){
			$(h2).removeClass("hover");
		}, 300);
	});
}


// RESET HOVER EFFECTS
function resetHoverEffects() {
	$(".boxLink .bgPic img").removeClass("animate");
}


// CONTACT FORM
function contactForm() {
	// HIDE DEFAULT FILE INPUT
	$('.fileInput').hide();

	if ($("html").attr("lang") == "de") {
		// INSERT NEW FILE INPUT
		$('<div class="niceFileInput"><span>Datei anhängen</span><button>durchsuchen</button></div>').insertAfter(".fileInput");
	} else {
		// INSERT NEW FILE INPUT
		$('<div class="niceFileInput"><span>Attach file</span><button>browse</button></div>').insertAfter(".fileInput");
	}

	// BIND CLICK TO NEW FILE INPUT
	$(document).on( "click", ".niceFileInput", function(event) {
		$(".niceFileInput").removeClass("selected");
        $('input[type="file"].fileInput').click();
    });

    // ON FILE SELECT
	$(".fileInput").change(function (){
		var fileName = $(this).val();

		$(".niceFileInput").addClass("selected");


		if ($("html").attr("lang") == "de") {
			// INSERT NEW FILE INPUT
			$(".niceFileInput span").text("Eine Datei wurde angehangen");
		} else {
			// INSERT NEW FILE INPUT
			$(".niceFileInput span").text("A file was attached");
		}
	});
}


// CUSTOM GOOGLE MAP
function googleMap() {
	google.maps.event.addDomListener(window, 'load', init);

	function init() {
		var mapOptions = {
			// DISABLE UI
			disableDefaultUI: true,

			// ZOOM LEVEL
			zoom: 9,

			// MAP LOCATION
			center: new google.maps.LatLng(49.706590, 8.779679),

			// MAP STYLES
			styles: [
					{"featureType":"road","elementType":"labels.text.fill","stylers":[{"color":"#696969"}]},
					{"featureType":"road","elementType":"labels.icon","stylers":[{"visibility":"off"}]},{},
					{"featureType":"road.highway","elementType":"geometry.stroke","stylers":[{"visibility":"on"},{"color":"#b3b3b3"}]},
					{"featureType":"road.highway","elementType":"geometry.fill","stylers":[{"color":"#eaeaea"}]},
					{"featureType":"road.local","elementType":"geometry.fill","stylers":[{"visibility":"on"},{"color":"#eaeaea"}]},
					{"featureType":"road.local","elementType":"geometry.stroke","stylers":[{"color":"#d7d7d7"}]},
					{"featureType":"road.arterial","elementType":"geometry.stroke","stylers":[{"color":"#c7ccd0"}]},
					{"featureType":"road.arterial","elementType":"geometry.fill","stylers":[{"color":"#eaeaea"}]},

					{"featureType":"poi","elementType":"geometry.fill","stylers":[{"visibility":"on"},{"color":"#ebebeb"}]},
					{"featureType":"poi","elementType":"labels.icon","stylers":[{"visibility":"on"},{"saturation":"-100"}]},
					{"featureType":"poi","elementType":"labels","stylers":[{"visibility":"on"},{"saturation":"-100"}]},

					{"featureType":"administrative","elementType":"geometry","stylers":[{"color":"#b7bfc6"}]},
					{"featureType":"administrative.province","elementType":"geometry.stroke","stylers":[{"visibility":"off"}]},
					{"featureType":"administrative","elementType":"labels.text.fill","stylers":[{"visibility":"on"},{"color":"#737373"}]},

					{"featureType":"landscape","elementType":"landscape.man_made","stylers":[{"visibility":"off"}]},
					{"featureType":"landscape","elementType":"geometry.fill","stylers":[{"visibility":"on"},{"color":"#e5e6e7"}]},

					{"featureType":"water","elementType":"geometry.fill","stylers":[{"color":"#d3d3d3"}]},
					{"featureType":"transit","stylers":[{"color":"#d3d3d3"},{"visibility":"off"}]},
			]
		};

		// CREATE DOM
		var mapElement = document.getElementById('map');
		var map = new google.maps.Map(mapElement, mapOptions);
		var iconPath = "#";

		// GET SVG ICON
		if ($("body").hasClass("ie8")) {
			iconPath = "files/kopp/img/icons/map-marker.png";
		} else {
			iconPath = "files/kopp/img/icons/map-marker.svg";
		}

		// SET MARKER
		var contentString = '<strong>KOPP Schleiftechnik</strong><br>Am Raupenstein 21<br> 64678 Lindenfels';
		var infowindow = new google.maps.InfoWindow({
			content: contentString
		});

		var marker = new google.maps.Marker({
			position: new google.maps.LatLng(49.706679, 8.776963),
			map: map,
			icon: iconPath,
			title: 'KOPP Schleiftechnik',
			id: "mapMarker"
		});

		marker.addListener('click', function() {
			infowindow.open(map, marker);
		});
	}
}


// FIXES
function fixes() {
	// REMOVE LINE BREAKS ON MOBILE
	if ($(window).width() < 1400) {
		$("p").each(function() {
			if($(this).html().replace(/\s|&nbsp;/g, '').length == 0)
				$(this).hide();
		});
	} else {
		$("p").each(function() {
			if($(this).html().replace(/\s|&nbsp;/g, '').length == 0)
				$(this).show();
		});
	}

	// REPLACE SVG WITH PNG IF IE8
	if ($("body").hasClass("ie8")) {
		$('img[src$="svg"]').attr('src', function() {  
			return $(this).attr('src').replace('.svg', '.png');
		})
	}

	// SLIDER ARROWS
	$(".slider-prev, .slider-next").css({
		"top": - $(".slideShow .content-slider").height() / 2 - 50
	});


	// VHM SITE
	if ($(window).width() < 1200) {
		$(".vhmCat").find(".span_3_of_12").each(function(){
			$(this).removeClass("span_3_of_12");
			$(this).addClass("span_5_of_12");
		});

		$(".vhmCat").find(".span_9_of_12").each(function(){
			$(this).removeClass("span_9_of_12");
			$(this).addClass("span_7_of_12");
		});

		$(".vhmPros").find(".span_4_of_12").each(function(){
			$(this).removeClass("span_4_of_12");
			$(this).addClass("span_6_of_12");
		});

		$(".vhmPros").find(".span_8_of_12").each(function(){
			$(this).removeClass("span_8_of_12");
			$(this).addClass("span_6_of_12");
		});

		$("#bohrerPros, #fraserPros").removeClass("span_8_of_12").addClass("span_11_of_12").addClass("first");
	}
}


// SET HEIGHTS
function setHeight() {
	if ($(window).width() > 500) {
		$(".equalLeft").each(function(){
			$(this).css({
				"height": $(this).prev().height(),
				"padding-top": ($(this).height() - $("this").find("h2, .headline").height()) / 6.5
			});
		});

		$(".equalRight").each(function(){
			$(this).css({
				"height": $(this).next().height() - 8,
				"padding-top": ($(this).height() - $("this").find("h2, .headline").height()) / 6.5
			});
		});
	} else {
		$(".equalRight, .equalLeft").css("height", "auto");
	}
}



// SET VIDEO LOOPS
function videos() {
	$(".loopVideo video").each(function(){
		$(this).attr("loop", "loop");
		$(this).removeAttr("controls");
	});
}



// NEWS FUNCTION
function newsList() {

	$(".layout_latest").each(function() {
		var link = $(this).find(".more a").attr("href");
		$(this).find("img").wrap( "<a href='#' class='newsLink'></div>" );
		$(this).find(".newsLink").attr("href", link);
	});

}


// DOWNLOAD FUNCTION
/*
function downloadList(){

	$(".downloadBox h2").each(function(){
		$(this).attr("title", $(this).text());
	});

	$(".ce_previewdownload").each(function() {
		var src = $(this).find(".preview_image").attr("src");

		$(this).find(".preview_image").addClass("deleteSoon");

		$('<div class="preview_image"></div>').insertAfter($(this).find(".preview_image"));
		$(this).find(".deleteSoon").remove();
		$(this).find(".preview_image").css('background-image', 'url("' + src + '")');
	});

	$(".ce_previewdownload a").removeAttr("onclick").attr("target", "_blank");

}

*/

// SLIDESHOW FUNCTION
function slideShow() {
	$(".slideShowAppend").appendTo(".slideShowIntroBox");
}


// POSITION BREADCRUMB NAV FUNCTION
function breadCrumb() {
	$(".mod_breadcrumb").insertAfter(".firstContent");
	$(".mod_breadcrumb").show();
}