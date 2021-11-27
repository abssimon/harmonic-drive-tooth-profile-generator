(function (cjs, an) {

var p; // shortcut to reference prototypes
var lib={};var ss={};var img={};
lib.ssMetadata = [
		{name:"Reinbold_final_4_atlas_1", frames: [[1990,0,44,62],[1880,137,46,62],[1979,194,39,62],[1990,64,44,62],[1880,201,46,62],[1988,128,41,64],[1520,306,444,121],[1936,67,50,65],[1880,0,54,68],[1972,258,37,65],[1930,134,47,65],[2026,356,13,65],[1936,0,52,65],[1928,201,42,65],[1880,70,48,65],[2020,194,21,43],[1880,265,46,34],[2011,258,30,49],[1997,356,27,34],[2031,128,6,46],[1979,134,6,47],[1966,374,23,34],[1966,325,29,47],[1991,392,23,34],[1928,268,28,34],[2011,309,32,45],[938,306,580,209],[938,0,940,304],[0,485,796,106],[0,0,936,360],[1520,429,231,199],[0,362,804,121]]}
];


(lib.AnMovieClip = function(){
	this.currentSoundStreamInMovieclip;
	this.actionFrames = [];
	this.soundStreamDuration = new Map();
	this.streamSoundSymbolsList = [];

	this.gotoAndPlayForStreamSoundSync = function(positionOrLabel){
		cjs.MovieClip.prototype.gotoAndPlay.call(this,positionOrLabel);
	}
	this.gotoAndPlay = function(positionOrLabel){
		this.clearAllSoundStreams();
		this.startStreamSoundsForTargetedFrame(positionOrLabel);
		cjs.MovieClip.prototype.gotoAndPlay.call(this,positionOrLabel);
	}
	this.play = function(){
		this.clearAllSoundStreams();
		this.startStreamSoundsForTargetedFrame(this.currentFrame);
		cjs.MovieClip.prototype.play.call(this);
	}
	this.gotoAndStop = function(positionOrLabel){
		cjs.MovieClip.prototype.gotoAndStop.call(this,positionOrLabel);
		this.clearAllSoundStreams();
	}
	this.stop = function(){
		cjs.MovieClip.prototype.stop.call(this);
		this.clearAllSoundStreams();
	}
	this.startStreamSoundsForTargetedFrame = function(targetFrame){
		for(var index=0; index<this.streamSoundSymbolsList.length; index++){
			if(index <= targetFrame && this.streamSoundSymbolsList[index] != undefined){
				for(var i=0; i<this.streamSoundSymbolsList[index].length; i++){
					var sound = this.streamSoundSymbolsList[index][i];
					if(sound.endFrame > targetFrame){
						var targetPosition = Math.abs((((targetFrame - sound.startFrame)/lib.properties.fps) * 1000));
						var instance = playSound(sound.id);
						var remainingLoop = 0;
						if(sound.offset){
							targetPosition = targetPosition + sound.offset;
						}
						else if(sound.loop > 1){
							var loop = targetPosition /instance.duration;
							remainingLoop = Math.floor(sound.loop - loop);
							if(targetPosition == 0){ remainingLoop -= 1; }
							targetPosition = targetPosition % instance.duration;
						}
						instance.loop = remainingLoop;
						instance.position = Math.round(targetPosition);
						this.InsertIntoSoundStreamData(instance, sound.startFrame, sound.endFrame, sound.loop , sound.offset);
					}
				}
			}
		}
	}
	this.InsertIntoSoundStreamData = function(soundInstance, startIndex, endIndex, loopValue, offsetValue){ 
 		this.soundStreamDuration.set({instance:soundInstance}, {start: startIndex, end:endIndex, loop:loopValue, offset:offsetValue});
	}
	this.clearAllSoundStreams = function(){
		var keys = this.soundStreamDuration.keys();
		for(var i = 0;i<this.soundStreamDuration.size; i++){
			var key = keys.next().value;
			key.instance.stop();
		}
 		this.soundStreamDuration.clear();
		this.currentSoundStreamInMovieclip = undefined;
	}
	this.stopSoundStreams = function(currentFrame){
		if(this.soundStreamDuration.size > 0){
			var keys = this.soundStreamDuration.keys();
			for(var i = 0; i< this.soundStreamDuration.size ; i++){
				var key = keys.next().value; 
				var value = this.soundStreamDuration.get(key);
				if((value.end) == currentFrame){
					key.instance.stop();
					if(this.currentSoundStreamInMovieclip == key) { this.currentSoundStreamInMovieclip = undefined; }
					this.soundStreamDuration.delete(key);
				}
			}
		}
	}

	this.computeCurrentSoundStreamInstance = function(currentFrame){
		if(this.currentSoundStreamInMovieclip == undefined){
			if(this.soundStreamDuration.size > 0){
				var keys = this.soundStreamDuration.keys();
				var maxDuration = 0;
				for(var i=0;i<this.soundStreamDuration.size;i++){
					var key = keys.next().value;
					var value = this.soundStreamDuration.get(key);
					if(value.end > maxDuration){
						maxDuration = value.end;
						this.currentSoundStreamInMovieclip = key;
					}
				}
			}
		}
	}
	this.getDesiredFrame = function(currentFrame, calculatedDesiredFrame){
		for(var frameIndex in this.actionFrames){
			if((frameIndex > currentFrame) && (frameIndex < calculatedDesiredFrame)){
				return frameIndex;
			}
		}
		return calculatedDesiredFrame;
	}

	this.syncStreamSounds = function(){
		this.stopSoundStreams(this.currentFrame);
		this.computeCurrentSoundStreamInstance(this.currentFrame);
		if(this.currentSoundStreamInMovieclip != undefined){
			var soundInstance = this.currentSoundStreamInMovieclip.instance;
			if(soundInstance.position != 0){
				var soundValue = this.soundStreamDuration.get(this.currentSoundStreamInMovieclip);
				var soundPosition = (soundValue.offset?(soundInstance.position - soundValue.offset): soundInstance.position);
				var calculatedDesiredFrame = (soundValue.start)+((soundPosition/1000) * lib.properties.fps);
				if(soundValue.loop > 1){
					calculatedDesiredFrame +=(((((soundValue.loop - soundInstance.loop -1)*soundInstance.duration)) / 1000) * lib.properties.fps);
				}
				calculatedDesiredFrame = Math.floor(calculatedDesiredFrame);
				var deltaFrame = calculatedDesiredFrame - this.currentFrame;
				if(deltaFrame >= 2){
					this.gotoAndPlayForStreamSoundSync(this.getDesiredFrame(this.currentFrame,calculatedDesiredFrame));
				}
			}
		}
	}
}).prototype = p = new cjs.MovieClip();
// symbols:



(lib.CachedBmp_97 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(0);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_95 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(1);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_96 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(2);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_92 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(3);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_91 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(4);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_90 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(5);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_98 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(6);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_89 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(7);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_87 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(8);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_88 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(9);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_86 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(10);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_84 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(11);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_85 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(12);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_83 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(13);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_82 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(14);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_81 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(15);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_79 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(16);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_78 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(17);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_77 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(18);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_80 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(19);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_75 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(20);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_74 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(21);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_73 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(22);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_72 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(23);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_71 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(24);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_70 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(25);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_69 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(26);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_68 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(27);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_65 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(28);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_67 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(29);
}).prototype = p = new cjs.Sprite();



(lib.azrk = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(30);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_99 = function() {
	this.initialize(ss["Reinbold_final_4_atlas_1"]);
	this.gotoAndStop(31);
}).prototype = p = new cjs.Sprite();



(lib.CachedBmp_66 = function() {
	this.initialize(img.CachedBmp_66);
}).prototype = p = new cjs.Bitmap();
p.nominalBounds = new cjs.Rectangle(0,0,4252,4252);// helper functions:

function mc_symbol_clone() {
	var clone = this._cloneProps(new this.constructor(this.mode, this.startPosition, this.loop));
	clone.gotoAndStop(this.currentFrame);
	clone.paused = this.paused;
	clone.framerate = this.framerate;
	return clone;
}

function getMCSymbolPrototype(symbol, nominalBounds, frameBounds) {
	var prototype = cjs.extend(symbol, cjs.MovieClip);
	prototype.clone = mc_symbol_clone;
	prototype.nominalBounds = nominalBounds;
	prototype.frameBounds = frameBounds;
	return prototype;
	}


(lib.ZERKLEINERUNG = function(mode,startPosition,loop) {
	this.initialize(mode,startPosition,loop,{});

	// Layer_1
	this.instance = new lib.azrk();
	this.instance.setTransform(-47,-28,0.8995,0.8995);

	this.timeline.addTween(cjs.Tween.get(this.instance).wait(1));

	this._renderFirstFrame();

}).prototype = getMCSymbolPrototype(lib.ZERKLEINERUNG, new cjs.Rectangle(-47,-28,207.8,179), null);


(lib.www = function(mode,startPosition,loop) {
	this.initialize(mode,startPosition,loop,{});

	// Ebene_1
	this.instance = new lib.CachedBmp_98();
	this.instance.setTransform(110.4,0,0.5,0.5);

	this.timeline.addTween(cjs.Tween.get(this.instance).wait(1));

	this._renderFirstFrame();

}).prototype = getMCSymbolPrototype(lib.www, new cjs.Rectangle(110.4,0,221.99999999999997,60.5), null);


(lib.wirstellen = function(mode,startPosition,loop) {
	this.initialize(mode,startPosition,loop,{});

	// FlashAICB
	this.instance = new lib.CachedBmp_97();
	this.instance.setTransform(280.65,82.1,0.5,0.5);

	this.instance_1 = new lib.CachedBmp_96();
	this.instance_1.setTransform(255.7,82.1,0.5,0.5);

	this.instance_2 = new lib.CachedBmp_95();
	this.instance_2.setTransform(226.65,82.1,0.5,0.5);

	this.instance_3 = new lib.CachedBmp_95();
	this.instance_3.setTransform(197.6,82.1,0.5,0.5);

	this.instance_4 = new lib.CachedBmp_96();
	this.instance_4.setTransform(172.65,82.1,0.5,0.5);

	this.instance_5 = new lib.CachedBmp_92();
	this.instance_5.setTransform(144.8,82.1,0.5,0.5);

	this.instance_6 = new lib.CachedBmp_91();
	this.instance_6.setTransform(114.3,82.1,0.5,0.5);

	this.instance_7 = new lib.CachedBmp_90();
	this.instance_7.setTransform(88.1,81.55,0.5,0.5);

	this.instance_8 = new lib.CachedBmp_89();
	this.instance_8.setTransform(275.75,31.5,0.5,0.5);

	this.instance_9 = new lib.CachedBmp_88();
	this.instance_9.setTransform(253.05,31.5,0.5,0.5);

	this.instance_10 = new lib.CachedBmp_87();
	this.instance_10.setTransform(220.35,30.95,0.5,0.5);

	this.instance_11 = new lib.CachedBmp_86();
	this.instance_11.setTransform(193,31.5,0.5,0.5);

	this.instance_12 = new lib.CachedBmp_85();
	this.instance_12.setTransform(159.75,31.5,0.5,0.5);

	this.instance_13 = new lib.CachedBmp_84();
	this.instance_13.setTransform(146,31.5,0.5,0.5);

	this.instance_14 = new lib.CachedBmp_83();
	this.instance_14.setTransform(119.4,31.5,0.5,0.5);

	this.instance_15 = new lib.CachedBmp_82();
	this.instance_15.setTransform(89.6,31.5,0.5,0.5);

	this.instance_16 = new lib.CachedBmp_81();
	this.instance_16.setTransform(287.15,-6.2,0.5,0.5);

	this.instance_17 = new lib.CachedBmp_80();
	this.instance_17.setTransform(279.2,-7.85,0.5,0.5);

	this.instance_18 = new lib.CachedBmp_79();
	this.instance_18.setTransform(249.55,-1.75,0.5,0.5);

	this.instance_19 = new lib.CachedBmp_78();
	this.instance_19.setTransform(220,-1.65,0.5,0.5);

	this.instance_20 = new lib.CachedBmp_77();
	this.instance_20.setTransform(200.7,-1.75,0.5,0.5);

	this.instance_21 = new lib.CachedBmp_80();
	this.instance_21.setTransform(190.95,-7.85,0.5,0.5);

	this.instance_22 = new lib.CachedBmp_75();
	this.instance_22.setTransform(180.9,-8.3,0.5,0.5);

	this.instance_23 = new lib.CachedBmp_74();
	this.instance_23.setTransform(163.35,-1.75,0.5,0.5);

	this.instance_24 = new lib.CachedBmp_73();
	this.instance_24.setTransform(144.55,-1.35,0.5,0.5);

	this.instance_25 = new lib.CachedBmp_72();
	this.instance_25.setTransform(128.65,-1.75,0.5,0.5);

	this.instance_26 = new lib.CachedBmp_71();
	this.instance_26.setTransform(109.75,-1.75,0.5,0.5);

	this.instance_27 = new lib.CachedBmp_70();
	this.instance_27.setTransform(88.8,-7.55,0.5,0.5);

	this.timeline.addTween(cjs.Tween.get({}).to({state:[{t:this.instance_27},{t:this.instance_26},{t:this.instance_25},{t:this.instance_24},{t:this.instance_23},{t:this.instance_22},{t:this.instance_21},{t:this.instance_20},{t:this.instance_19},{t:this.instance_18},{t:this.instance_17},{t:this.instance_16},{t:this.instance_15},{t:this.instance_14},{t:this.instance_13},{t:this.instance_12},{t:this.instance_11},{t:this.instance_10},{t:this.instance_9},{t:this.instance_8},{t:this.instance_7},{t:this.instance_6},{t:this.instance_5},{t:this.instance_4},{t:this.instance_3},{t:this.instance_2},{t:this.instance_1},{t:this.instance}]}).wait(1));

	this._renderFirstFrame();

}).prototype = getMCSymbolPrototype(lib.wirstellen, new cjs.Rectangle(88.1,-8.3,214.6,121.89999999999999), null);


(lib.Text_Materialen = function(mode,startPosition,loop) {
	this.initialize(mode,startPosition,loop,{});

	// Ebene_1
	this.instance = new lib.CachedBmp_69();
	this.instance.setTransform(66.85,16,0.5,0.5);

	this.timeline.addTween(cjs.Tween.get(this.instance).wait(1));

	this._renderFirstFrame();

}).prototype = getMCSymbolPrototype(lib.Text_Materialen, new cjs.Rectangle(66.9,16,290,104.5), null);


(lib.rot = function(mode,startPosition,loop) {
	this.initialize(mode,startPosition,loop,{});

	// Ebene_1
	this.instance = new lib.CachedBmp_68();
	this.instance.setTransform(0,0,0.5,0.5);

	this.timeline.addTween(cjs.Tween.get(this.instance).wait(1));

	this._renderFirstFrame();

}).prototype = getMCSymbolPrototype(lib.rot, new cjs.Rectangle(0,0,470,152), null);


(lib.link = function(mode,startPosition,loop) {
	this.initialize(mode,startPosition,loop,{});

	// Ebene_1
	this.instance = new lib.CachedBmp_67();
	this.instance.setTransform(0,0,0.5,0.5);
	this.instance._off = true;

	this.timeline.addTween(cjs.Tween.get(this.instance).wait(3).to({_off:false},0).wait(1));

	this._renderFirstFrame();

}).prototype = p = new cjs.MovieClip();
p.nominalBounds = new cjs.Rectangle(0,0,468,180);


(lib.Earth = function(mode,startPosition,loop) {
	this.initialize(mode,startPosition,loop,{});

	// Ebene_1
	this.instance = new lib.CachedBmp_66();
	this.instance.setTransform(0,0,0.5,0.5);

	this.timeline.addTween(cjs.Tween.get(this.instance).wait(1));

	this._renderFirstFrame();

}).prototype = getMCSymbolPrototype(lib.Earth, new cjs.Rectangle(0,0,2126,2126), null);


// stage content:
(lib.Reibold_HTML5Canvas = function(mode,startPosition,loop) {
	this.initialize(mode,startPosition,loop,{});

	this.actionFrames = [0,294];
	// timeline functions:
	this.frame_0 = function() {
		this.clearAllSoundStreams();
		 
		/* Klicken, um zu Webseite zu gehen
		Durch Klicken auf die angegebene Symbolinstanz wird die URL in einem neuen Browserfenster geladen.
		
		Anweisungen:
		1. Ersetzen Sie http://www.adobe.com durch die gewünschte URL-Adresse.
		      Lassen Sie die Anführungszeichen ("") stehen.
		*/
		
		this.button_1.addEventListener("click", fl_ClickToGoToWebPage);
		
		function fl_ClickToGoToWebPage() {
			window.open("http://reinbold.de/", "_blank");
		}
	}
	this.frame_294 = function() {
		/* Klicken, um zu Webseite zu gehen
		Durch Klicken auf die angegebene Symbolinstanz wird die URL in einem neuen Browserfenster geladen.
		
		Anweisungen:
		1. Ersetzen Sie http://www.adobe.com durch die gewünschte URL-Adresse.
		      Lassen Sie die Anführungszeichen ("") stehen.
		*/
		
		this.button_1.addEventListener("click", fl_ClickToGoToWebPage);
		
		function fl_ClickToGoToWebPage() {
			window.open("https://www.reinbold-entsorgungstechnik.com/", "_blank");
		}
	}

	// actions tween:
	this.timeline.addTween(cjs.Tween.get(this).call(this.frame_0).wait(294).call(this.frame_294).wait(3));

	// Ebene_2
	this.button_1 = new lib.link();
	this.button_1.name = "button_1";
	this.button_1.setTransform(117,30,1,1,0,0,0,117.4,30.4);
	new cjs.ButtonHelper(this.button_1, 0, 1, 2, false, new lib.link(), 3);

	this.timeline.addTween(cjs.Tween.get(this.button_1).to({_off:true},1).wait(1).to({_off:false},0).wait(293).to({_off:true},1).wait(1));

	// www
	this.instance = new lib.www();
	this.instance.setTransform(233.95,126.3,1,1,0,0,0,221.3,26.4);
	this.instance.alpha = 0;
	this.instance._off = true;

	this.timeline.addTween(cjs.Tween.get(this.instance).wait(210).to({_off:false},0).to({alpha:1},9).wait(70).to({alpha:0},6).to({_off:true},1).wait(1));

	// Earth
	this.instance_1 = new lib.Earth();
	this.instance_1.setTransform(234.05,90.05,1,1,0,0,0,1063,1063);
	this.instance_1.alpha = 0;
	this.instance_1._off = true;

	this.timeline.addTween(cjs.Tween.get(this.instance_1).wait(174).to({_off:false},0).to({alpha:1},16).wait(3).to({scaleX:0.0239,scaleY:0.0239,x:233.6,y:65.45},20).to({_off:true},83).wait(1));

	// Text
	this.instance_2 = new lib.CachedBmp_99();
	this.instance_2.setTransform(33.1,-3.7,0.5,0.5);

	this.instance_3 = new lib.CachedBmp_65();
	this.instance_3.setTransform(16,-3.7,0.5,0.5);

	this.timeline.addTween(cjs.Tween.get({}).to({state:[]}).to({state:[{t:this.instance_2}]},80).to({state:[{t:this.instance_3}]},117).to({state:[]},2).wait(98));

	// Rot
	this.instance_4 = new lib.rot();
	this.instance_4.setTransform(233,239.05,1,1,0,0,0,235,76);
	this.instance_4.alpha = 0;

	this.timeline.addTween(cjs.Tween.get(this.instance_4).to({alpha:1},5).wait(69).to({y:-26.95},6).to({_off:true},117).wait(100));

	// Text_Materialen
	this.instance_5 = new lib.Text_Materialen();
	this.instance_5.setTransform(226.6,155.05,1,1,0,0,0,203.6,111);
	this.instance_5.alpha = 0;
	this.instance_5._off = true;

	this.timeline.addTween(cjs.Tween.get(this.instance_5).wait(80).to({_off:false},0).to({alpha:1},9).wait(101).to({y:-51.95},0).to({_off:true},2).wait(105));

	// Reinbold
	this.instance_6 = new lib.wirstellen();
	this.instance_6.setTransform(141.85,28.6);
	this.instance_6.alpha = 0;

	this.timeline.addTween(cjs.Tween.get(this.instance_6).to({alpha:1},5).wait(69).to({alpha:0},5).to({_off:true},3).wait(215));

	// ZERKLEINERUNG
	this.instance_7 = new lib.ZERKLEINERUNG();
	this.instance_7.setTransform(-161.2,29);

	this.timeline.addTween(cjs.Tween.get(this.instance_7).to({x:46.6},6).wait(69).to({alpha:0},6).to({_off:true},1).wait(215));

	this._renderFirstFrame();

}).prototype = p = new lib.AnMovieClip();
p.nominalBounds = new cjs.Rectangle(-594.9,-882.9,1892,2036);
// library properties:
lib.properties = {
	id: 'D09876C5CAA54D22932013A8C4957613',
	width: 468,
	height: 180,
	fps: 24,
	color: "#E8E8E8",
	opacity: 1.00,
	manifest: [
		{src:"images/CachedBmp_66.png", id:"CachedBmp_66"},
		{src:"images/Reinbold_final_4_atlas_1.png", id:"Reinbold_final_4_atlas_1"}
	],
	preloads: []
};



// bootstrap callback support:

(lib.Stage = function(canvas) {
	createjs.Stage.call(this, canvas);
}).prototype = p = new createjs.Stage();

p.setAutoPlay = function(autoPlay) {
	this.tickEnabled = autoPlay;
}
p.play = function() { this.tickEnabled = true; this.getChildAt(0).gotoAndPlay(this.getTimelinePosition()) }
p.stop = function(ms) { if(ms) this.seek(ms); this.tickEnabled = false; }
p.seek = function(ms) { this.tickEnabled = true; this.getChildAt(0).gotoAndStop(lib.properties.fps * ms / 1000); }
p.getDuration = function() { return this.getChildAt(0).totalFrames / lib.properties.fps * 1000; }

p.getTimelinePosition = function() { return this.getChildAt(0).currentFrame / lib.properties.fps * 1000; }

an.bootcompsLoaded = an.bootcompsLoaded || [];
if(!an.bootstrapListeners) {
	an.bootstrapListeners=[];
}

an.bootstrapCallback=function(fnCallback) {
	an.bootstrapListeners.push(fnCallback);
	if(an.bootcompsLoaded.length > 0) {
		for(var i=0; i<an.bootcompsLoaded.length; ++i) {
			fnCallback(an.bootcompsLoaded[i]);
		}
	}
};

an.compositions = an.compositions || {};
an.compositions['D09876C5CAA54D22932013A8C4957613'] = {
	getStage: function() { return exportRoot.stage; },
	getLibrary: function() { return lib; },
	getSpriteSheet: function() { return ss; },
	getImages: function() { return img; }
};

an.compositionLoaded = function(id) {
	an.bootcompsLoaded.push(id);
	for(var j=0; j<an.bootstrapListeners.length; j++) {
		an.bootstrapListeners[j](id);
	}
}

an.getComposition = function(id) {
	return an.compositions[id];
}


an.makeResponsive = function(isResp, respDim, isScale, scaleType, domContainers) {		
	var lastW, lastH, lastS=1;		
	window.addEventListener('resize', resizeCanvas);		
	resizeCanvas();		
	function resizeCanvas() {			
		var w = lib.properties.width, h = lib.properties.height;			
		var iw = window.innerWidth, ih=window.innerHeight;			
		var pRatio = window.devicePixelRatio || 1, xRatio=iw/w, yRatio=ih/h, sRatio=1;			
		if(isResp) {                
			if((respDim=='width'&&lastW==iw) || (respDim=='height'&&lastH==ih)) {                    
				sRatio = lastS;                
			}				
			else if(!isScale) {					
				if(iw<w || ih<h)						
					sRatio = Math.min(xRatio, yRatio);				
			}				
			else if(scaleType==1) {					
				sRatio = Math.min(xRatio, yRatio);				
			}				
			else if(scaleType==2) {					
				sRatio = Math.max(xRatio, yRatio);				
			}			
		}			
		domContainers[0].width = w * pRatio * sRatio;			
		domContainers[0].height = h * pRatio * sRatio;			
		domContainers.forEach(function(container) {				
			container.style.width = w * sRatio + 'px';				
			container.style.height = h * sRatio + 'px';			
		});			
		stage.scaleX = pRatio*sRatio;			
		stage.scaleY = pRatio*sRatio;			
		lastW = iw; lastH = ih; lastS = sRatio;            
		stage.tickOnUpdate = false;            
		stage.update();            
		stage.tickOnUpdate = true;		
	}
}
an.handleSoundStreamOnTick = function(event) {
	if(!event.paused){
		var stageChild = stage.getChildAt(0);
		if(!stageChild.paused){
			stageChild.syncStreamSounds();
		}
	}
}


})(createjs = createjs||{}, AdobeAn = AdobeAn||{});
var createjs, AdobeAn;