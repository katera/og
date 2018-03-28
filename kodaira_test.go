package og_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/katera/og"
)

const (
	mockHtml = `
	<!DOCTYPE html>
	<!--[if (gt IE 9)|!(IE)]> <!--> <html lang="en" class="no-js section-technology format-extra-long tone-feature app-article page-theme-standard has-comments has-cover-media no-ribbon type-size-small" itemid="https://www.nytimes.com/2018/03/21/technology/facebook-zucktown-willow-village.html" itemtype="http://schema.org/NewsArticle"  itemscope xmlns:og="http://opengraphprotocol.org/schema/"> <!--<![endif]-->
	<!--[if IE 9]> <html lang="en" class="no-js ie9 lt-ie10 section-technology format-extra-long tone-feature app-article page-theme-standard has-comments has-cover-media no-ribbon type-size-small" xmlns:og="http://opengraphprotocol.org/schema/"> <![endif]-->
	<!--[if IE 8]> <html lang="en" class="no-js ie8 lt-ie10 lt-ie9 section-technology format-extra-long tone-feature app-article page-theme-standard has-comments has-cover-media no-ribbon type-size-small" xmlns:og="http://opengraphprotocol.org/schema/"> <![endif]-->
	<!--[if (lt IE 8)]> <html lang="en" class="no-js lt-ie10 lt-ie9 lt-ie8 section-technology format-extra-long tone-feature app-article page-theme-standard has-comments has-cover-media no-ribbon type-size-small" xmlns:og="http://opengraphprotocol.org/schema/"> <![endif]-->
	<head>
		<title>Welcome to Zucktown. Where Everything Is Just Zucky. - The New York Times</title>
			<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
		<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1" />
	<link rel="shortcut icon" href="https://static01.nyt.com/favicon.ico" />
	<link rel="apple-touch-icon-precomposed" sizes="144×144" href="https://static01.nyt.com/images/icons/ios-ipad-144x144.png" />
	<link rel="apple-touch-icon-precomposed" sizes="114×114" href="https://static01.nyt.com/images/icons/ios-iphone-114x144.png" />
	<link rel="apple-touch-icon-precomposed" href="https://static01.nyt.com/images/icons/ios-default-homescreen-57x57.png" />
	<meta name="sourceApp" content="nyt-v5" />
	<meta id="applicationName" name="applicationName" content="article" />
	<meta id="foundation-build-id" name="foundation-build-id" content="" />
	<link rel="canonical" href="https://www.nytimes.com/2018/03/21/technology/facebook-zucktown-willow-village.html" />
	<link rel="alternate" media="only screen and (max-width: 640px)" href="http://mobile.nytimes.com/2018/03/21/technology/facebook-zucktown-willow-village.html" />
	<link rel="alternate" media="handheld" href="http://mobile.nytimes.com/2018/03/21/technology/facebook-zucktown-willow-village.html" />
	<link rel="alternate" hreflang="es-LA" href="https://www.nytimes.com/es/2018/03/24/facebook-google-zucktown-alphabet/?" />
	<link rel="alternate" hreflang="en" href="https://www.nytimes.com/2018/03/21/technology/facebook-zucktown-willow-village.html" />
	<meta property="al:android:url" content="nytimes://reader/id/100000005735109" />
	<meta property="al:android:package" content="com.nytimes.android" />
	<meta property="al:android:app_name" content="NYTimes" />
	<meta name="twitter:app:name:googleplay" content="NYTimes" />
	<meta name="twitter:app:id:googleplay" content="com.nytimes.android" />
	<meta name="twitter:app:url:googleplay" content="nytimes://reader/id/100000005735109" />
	<link rel="alternate" href="android-app://com.nytimes.android/nytimes/reader/id/100000005735109" />
	<meta property="al:iphone:url" content="nytimes://www.nytimes.com/2018/03/21/technology/facebook-zucktown-willow-village.html" />
	<meta property="al:iphone:app_store_id" content="284862083" />
	<meta property="al:iphone:app_name" content="NYTimes" />
	<meta property="al:ipad:url" content="nytimes://www.nytimes.com/2018/03/21/technology/facebook-zucktown-willow-village.html" />
	<meta property="al:ipad:app_store_id" content="357066198" />
	<meta property="al:ipad:app_name" content="NYTimes" />
	<meta name="robots" content="noarchive" />
	<meta itemprop="alternativeHeadline" name="hdl_p" content="Welcome to Zucktown" />
	<meta name="channels" content="" />
	<meta name="title" content="Welcome to Zucktown. Where Everything Is Just Zucky." />
	<meta itemprop="description" name="description" content="In Menlo Park, Calif., Facebook is building a real community and testing the proposition: Do people love tech companies so much they will live inside them?" />
	<meta name="genre" itemprop="genre" content="News" />
	<meta itemprop="identifier" name="articleid" content="100000005735109" />
	<meta itemprop="usageTerms" name="usageTerms" content="https://www.nytimes.com/content/help/rights/sale/terms-of-sale.html" />
	<meta itemprop="inLanguage" content="en-US" />
	<meta name="hdl" content="Welcome to Zucktown. Where Everything Is Just Zucky." />
	<meta name="col" content="" id="column-name" />
	<meta name="pdate" content="20180321" />
	<meta name="utime" content="20180326104551" />
	<meta name="ptime" content="20180321050012" />
	<meta name="DISPLAYDATE" content="March 21, 2018" />
	<meta name="dat" content="March 21, 2018" />
	<meta name="lp" content="In Menlo Park, Calif., Facebook is building a real community and testing the proposition: Do people love tech companies so much they will live inside them?" />
	<meta name="msapplication-starturl" content="http://www.nytimes.com" />
	<meta name="cre" content="The New York Times" />
	<meta name="slug" content="25zucktown" />
	<meta property="article:collection" content="https://static01.nyt.com/services/json/sectionfronts/technology/index.jsonp" />
	<meta name="sectionfront_jsonp" content="https://static01.nyt.com/services/json/sectionfronts/technology/index.jsonp" />
	<meta property="og:url" content="https://www.nytimes.com/2018/03/21/technology/facebook-zucktown-willow-village.html" />
	<meta property="og:type" content="article" />
	<meta property="og:site" content="nytimes.com" />
	<meta property="og:site_name" content="The New York Times" />
	<meta property="og:locale" content="en_US" />
	<meta property="og:title" content="Welcome to Zucktown. Where Everything Is Just Zucky." />
	<meta property="og:description" content="In Menlo Park, Calif., Facebook is building a real community and testing the proposition: Do people love tech companies so much they will live inside them?" />
	<meta property="article:published_time" itemprop="datePublished" content="2018-03-21T05:00:12-04:00" />
	<meta property="article:modified_time" itemprop="dateModified" content="2018-03-26T10:45:51-04:00" />
	<meta property="article:section" itemprop="articleSection" content="Technology" />
	<meta property="article:section-taxonomy-id" itemprop="articleSection" content="78FBAD45-31A9-4EC7-B172-7D62A2B9955E" />
	<meta property="article:section_url" content="https://www.nytimes.com/section/technology" />
	<meta property="article:top-level-section" content="technology" />
	<meta property="fb:app_id" content="9869919170" />
	<meta name="twitter:site" value="@nytimes" />
	<meta property="twitter:url" content="https://www.nytimes.com/2018/03/21/technology/facebook-zucktown-willow-village.html" />
	<meta property="twitter:title" content="Welcome to Zucktown. Where Everything Is Just Zucky." />
	<meta property="twitter:description" content="In Menlo Park, Calif., Facebook is building a real community and testing the proposition: Do people love tech companies so much they will live inside them?" />
	<meta name="author" content="David Streitfeld" />
	<meta name="tone" content="feature" id="article-tone" />
	<meta name="byl" content="By DAVID STREITFELD" />
	<meta name="PT" content="article" />
	<meta name="CG" content="technology" />
	<meta name="SCG" content="" />
	<meta name="PST" content="News" />
	<meta name="tom" content="News" />
	<meta name="edt" content="NewYork" />
	<meta property="og:image" content="https://static01.nyt.com/images/2018/03/20/business/00ZUCKTOWN-8/00ZUCKTOWN-8-facebookJumbo-v2.jpg" />
	<meta property="twitter:image:alt" content="Juan Salazar, a Facebook public policy manager, showed a model of the company&rsquo;s planned development to representatives of local businesses. &ldquo;Our goal is to strengthen the community,&rdquo; he said." />
	<meta property="twitter:image" content="https://static01.nyt.com/images/2018/03/20/business/00ZUCKTOWN-8/00ZUCKTOWN-8-videoSixteenByNineJumbo1600-v2.jpg" />
	<meta name="twitter:card" value="summary_large_image" />
	<meta property="article:author" content="https://www.nytimes.com/by/david-streitfeld" />
	<meta property="article:tag" content="Willow Village (Menlo Park, Calif)" />
	<meta name="geo" content="Willow Village (Menlo Park, Calif)" />
	<meta property="article:tag" content="Menlo Park (Calif)" />
	<meta name="geo" content="Menlo Park (Calif)" />
	<meta property="article:tag" content="Computers and the Internet" />
	<meta name="des" content="Computers and the Internet" />
	<meta property="article:tag" content="Area Planning and Renewal" />
	<meta name="des" content="Area Planning and Renewal" />
	<meta property="article:tag" content="Real Estate and Housing (Residential)" />
	<meta name="des" content="Real Estate and Housing (Residential)" />
	<meta property="article:tag" content="Social Media" />
	<meta name="des" content="Social Media" />
	<meta property="article:tag" content="Renting and Leasing (Real Estate)" />
	<meta name="des" content="Renting and Leasing (Real Estate)" />
	<meta property="article:tag" content="Real Estate (Commercial)" />
	<meta name="des" content="Real Estate (Commercial)" />
	<meta property="article:tag" content="Facebook Inc" />
	<meta name="org" content="Facebook Inc" />
	<meta property="article:tag" content="Alphabet Inc" />
	<meta name="org" content="Alphabet Inc" />
	<meta property="article:tag" content="Apple Inc" />
	<meta name="org" content="Apple Inc" />
	<meta property="article:tag" content="Chan Zuckerberg Initiative LLC" />
	<meta name="org" content="Chan Zuckerberg Initiative LLC" />
	<meta property="article:tag" content="Zuckerberg, Mark E" />
	<meta name="per" content="Zuckerberg, Mark E" />
	<meta property="article:tag" content="Mountain View (Calif)" />
	<meta name="geo" content="Mountain View (Calif)" />
	<meta property="article:tag" content="Cupertino (Calif)" />
	<meta name="geo" content="Cupertino (Calif)" />
	<meta name="keywords" content="Willow Village (Menlo Park  Calif),Menlo Park (Calif),Computers and the Internet,Area Planning and Renewal,Real Estate and Housing (Residential),Social Media,Renting and Leasing (Real Estate),Real Estate (Commercial),Facebook Inc,Alphabet Inc,Apple Inc,Chan Zuckerberg Initiative LLC,Zuckerberg  Mark E,Mountain View (Calif),Cupertino (Calif)" />
	<meta name="news_keywords" content="Willow Village Zucktown,Menlo Park CA,Computers and the Internet,Urban Planning,Real Estate and Housing,Social Media,Rent,Commercial Real Estate,Facebook,Alphabet,Apple,Chan Zuckerberg Initiative,Mark E Zuckerberg,Mount View CA,Cupertino" />
	<meta name="thumbnail_150" content="https://static01.nyt.com/images/2018/03/20/business/00ZUCKTOWN-8/00ZUCKTOWN-8-thumbLarge-v2.jpg" />
	<meta name="thumbnail_150_height" content="150" />
	<meta name="thumbnail_150_width" content="150" />
	<meta itemprop="thumbnailUrl" name="thumbnail" content="https://static01.nyt.com/images/2018/03/20/business/00ZUCKTOWN-8/00ZUCKTOWN-8-thumbStandard-v2.jpg" />
	<meta name="thumbnail_height" content="75" />
	<meta name="thumbnail_width" content="75" />
	<meta name="dfp-ad-unit-path" content="technology" />
	<meta name="dfp-amazon-enabled" content="false" />
	<link rel="alternate" type="application/json+oembed" href="https://www.nytimes.com/svc/oembed/json/?url=https%3A%2F%2Fwww.nytimes.com%2F2018%2F03%2F21%2Ftechnology%2Ffacebook-zucktown-willow-village.html" title="Welcome to Zucktown. Where Everything Is Just Zucky." />
	<meta property="fb:pages" content="5281959998" />
	<meta property="fb:pages" content="1002391179791389" />
	<meta property="fb:pages" content="105307012882667" />
	<meta property="fb:pages" content="123334484346296" />
	<meta property="fb:pages" content="142271415829971" />
	<meta property="fb:pages" content="1453472561556218" />
	<meta property="fb:pages" content="153491058028703" />
	<meta property="fb:pages" content="1545690368981813" />
	<meta property="fb:pages" content="1606969219532748" />
	<meta property="fb:pages" content="173455642668373" />
	<meta property="fb:pages" content="176999952314969" />
	<meta property="fb:pages" content="181548405298864" />
	<meta property="fb:pages" content="340711315527" />
	<meta property="fb:pages" content="5518834980" />
	<meta property="fb:pages" content="689907071131476" />
	<meta property="fb:pages" content="751965821517958" />
	<meta property="fb:pages" content="8254939527" />
	<meta property="fb:pages" content="98934992430" />
	<meta property="fb:pages" content="993603507345855" />
		
				<link id="legacy-zam5nzz" rel="stylesheet" type="text/css" href="https://typeface.nyt.com/css/zam5nzz.css" media="all" />
		<!--[if (gt IE 9)|!(IE)]> <!-->
			<link rel="stylesheet" type="text/css" media="screen" href="https://g1.nyt.com/assets/article/20180322-142545/css/article/story/styles.css" />
		<!--<![endif]-->
		<!--[if lte IE 9]>
			<link rel="stylesheet" type="text/css" media="screen" href="https://g1.nyt.com/assets/article/20180322-142545/css/article/story/styles-ie.css" />
		<![endif]-->
	<link rel="stylesheet" type="text/css" media="print" href="https://g1.nyt.com/assets/article/20180322-142545/css/article/story/styles-print.css" />
					<!--  begin abra  -->
	<script>
	var NYTD=NYTD||{};NYTD.Abra=function(t){"use strict";function n(t){var n=i[t];return n&&n[1]||null}function e(t,n){if(t){var e,o,r=n[0],i=n[1],u=0,c=0;if(1!==i.length||4294967296!==i[0])for(e=a(t+" "+r)>>>0,u=0,c=0;o=i[u++];)if(e<(c+=o[0]))return o}}function o(n,e,o,a){s+="subject="+n+"&test="+encodeURIComponent(e)+"&variant="+encodeURIComponent(o||0)+"&url="+encodeURIComponent(t.location.href)+"&instant=1&skipAugment=true\n",a&&f.push(a),c||(c=t.setTimeout(r,0))}function r(){var n=new t.XMLHttpRequest,e=f;n.withCredentials=!0,n.open("POST",u),n.onreadystatechange=function(){var t,o;if(4==n.readyState)for(t=200==n.status?null:new Error(n.statusText);o=e.shift();)o(t)},n.send(s),s="",f=[],c=null}function a(t){for(var n,e,o,r,a,i,u,c=0,s=0,f=[],l=[n=1732584193,e=4023233417,~n,~e,3285377520],h=[],p=t.length;s<=p;)h[s>>2]|=(s<p?t.charCodeAt(s):128)<<8*(3-s++%4);for(h[u=p+8>>2|15]=p<<3;c<=u;c+=16){for(n=l,s=0;s<80;n=[0|[(i=((t=n[0])<<5|t>>>27)+n[4]+(f[s]=s<16?~~h[c+s]:i<<1|i>>>31)+1518500249)+((e=n[1])&(o=n[2])|~e&(r=n[3])),a=i+(e^o^r)+341275144,i+(e&o|e&r|o&r)+882459459,a+1535694389][0|s++/20],t,e<<30|e>>>2,o,r])i=f[s-3]^f[s-8]^f[s-14]^f[s-16];for(s=5;s;)l[--s]=l[s]+n[s]|0}return l[0]}var i,u,c,s="",f=[];return n.init=function(n,r){var a,c,s,f,l,h=[],p=(t.document.cookie.match(/(^|;) *nyt-a=([^;]*)/)||[])[2],d=(t.document.cookie.match(/(^|;) *ab7=([^;]*)/)||[])[2],m=new RegExp("[?&]abra(=([^&#]*))"),g=m.exec(t.location.search);if(g&&g[2]&&(d=d?g[2]+"&"+d:g[2]),i)throw new Error("can't init twice");if(u=r,i={},d)for(d=decodeURIComponent(d).split("&"),a=0;a<d.length;a++)l=d[a].split("="),l[0]in i||(i[l[0]]=[,l[1],1],o("ab-alloc",l[0],l[1]),l[1]&&h.push(l[0]+"="+l[1]));for(a=0;a<n.length;a++)s=n[a],(c=s[0])in i||(f=e(p,s)||[],i[c]=f,f[1]&&h.push(c.replace(/[^\w-]/g)+"="+(""+f[1]).replace(/[^\w-]/g)),f[2]&&o("ab-alloc",c,f[1]));h.length&&t.document.documentElement.setAttribute("data-nyt-ab",h.join(" "))},n.reportExposure=function(n,e){var r=i[n];r&&r[2]?o("ab-expose",n,r[1],e):e&&t.setTimeout(function(){e(null)},0)},n}(this);
	</script>
	<script>NYTD.Abra.init([["www-story-reader-satisfaction",[[21474837,"Control",1],[21474836,"VariantA",1],[4252017623,null,0]]],["www-signup-favor-test",[[2147483648,"0",0],[2147483648,"1",0]]],["www-auto-newsletter",[[4294967296,"1",0]]],["www-STRIK-nl-regi",[[4294967296,"2",0]]],["abra_dfp",[[3865470567,"control",1],[429496729,"test",1]]],["EC_SampleTest",[[2147483648,"variantA",0],[2147483648,"variantB",0]]]], '//et.nytimes.com/')</script>
	<!--  end abra  -->
	
	<script>
	/* ABRA reporting for Vi rollout. Please don't hand-edit this minified snippet. <https://github.com/nytm/vi-rollout-abra-reporting-js> (6abdd06 @ Wed Feb  7 15:26:18 EST 2018) */ function reportViRolloutToABRA(t,e,o){"use strict";var i,a,n,r,s="";if(e&&(/^mobile\./.test(window.location.hostname)?(i=(document.cookie.match(/(?:^|;) *vi=([^;]*)/)||[,""])[1],a=(document.cookie.match(/(?:^|;) *nyt\.np\.vi=([^;]*)/)||[,""])[1],n=/^b1/.test(i)?"WP_ProjectVi&variant=vi":/^z1/.test(i)?"WP_ProjectVi&variant=0":/^a2/.test(i)?"WP_ProjectVi2&variant=hp-st":/^b2/.test(i)?"WP_ProjectVi2&variant=hp":/^c2/.test(i)?"WP_ProjectVi2&variant=st":/^z2/.test(i)?"WP_ProjectVi2&variant=0":/^1\|/.test(a)?"WP_ProjectVi&variant=vi":/^2\|/.test(a)?"WP_ProjectVi&variant=0":null):/^www\./.test(window.location.hostname)&&(i=(document.cookie.match(/(?:^|;) *vi_www_hp=([^;]*)/)||[,""])[1],n=/^a2/.test(i)?"WP_ProjectVi_www_hp&variant=hp-st":/^b2/.test(i)?"WP_ProjectVi_www_hp&variant=hp":/^c2/.test(i)?"WP_ProjectVi_www_hp&variant=st":/^d2/.test(i)?"WP_ProjectVi_www_hp&variant=hp-2":/^e2/.test(i)?"WP_ProjectVi_www_hp&variant=hp-serv":/^f2/.test(i)?"WP_ProjectVi_www_hp&variant=hp-orig":/^y2/.test(i)?"WP_ProjectVi_www_hp&variant=0-2":/^z2/.test(i)?"WP_ProjectVi_www_hp&variant=0":null),n&&(s+="subject=ab-alloc&test="+n+"&url="+encodeURIComponent(window.location.href)+"&instant=1&skipAugment=true\n")),o&&(/^mobile\./.test(window.location.hostname)?(i=(document.cookie.match(/(?:^|;) *vistorymobile=([^;]*)/)||[,""])[1],r=/^a/.test(i)?"WP_ProjectVi_Story_Mobile&variant=st":/^b/.test(i)?"WP_ProjectVi_Story_Mobile&variant=sth":/^c/.test(i)?"WP_ProjectVi_Story_Mobile&variant=mw":/^z/.test(i)?"WP_ProjectVi_Story_Mobile&variant=mwh":null):/^www\./.test(window.location.hostname)&&(i=(document.cookie.match(/(?:^|;) *vistory=([^;]*)/)||[,""])[1],r=/^a/.test(i)?"WP_ProjectVi_story_desktop&variant=sd":/^b/.test(i)?"WP_ProjectVi_story_desktop&variant=sdh":/^c/.test(i)?"WP_ProjectVi_story_desktop&variant=nyt5":/^z/.test(i)?"WP_ProjectVi_story_desktop&variant=nyt5h":null),r&&(s+="subject=ab-alloc&test="+r+"&url="+encodeURIComponent(window.location.href)+"&instant=1&skipAugment=true\n")),s){var _=new XMLHttpRequest;_.withCredentials=!0,_.open("POST",t),_.send(s)}}
	</script>
	<script>
	  reportViRolloutToABRA(
		'//et.nytimes.com/',
		false,
		true  );
	</script>
					<script id="page-config-data" type="text/json">
	{"pageconfig":{"ledeMediaSize":"full_bleed","keywords":["article-extra-long","has-embedded-interactive"],"collections":{"sections":["us","realestate","technology","economy","businessday"]}}}</script>
	<script id="display_overrides">
				["INCLUDE_TRANSPARENT_MASTHEAD_WITH_BLACK_ICONS","HIDE_RIBBON","HIDE_TOP_AD","HIDE_NAVIGATION-EDGE","HIDE_KICKER","HIDE_STORY-META_COMMUNITY_BUTTON"]    </script>
	
					<script type="text/javascript">
			(function() {
				var s = document.getElementsByTagName("script")[0];
				var mediaDotNet = 'https://contextual.media.net/bidexchange.js?cid=8CU2553YN&amp;https=1';
				var indexBid = 'https://js-sec.indexww.com/ht/p/183760-203795517182556.js';
				var timeout = 300;
	
				function loadScript(tagSrc) {
					if (tagSrc.substr(0, 4) !== 'http') {
						var isSSL = 'https:' == document.location.protocol;
						tagSrc = (isSSL ? 'https:' : '') + tagSrc;
					}
					var scriptTag = document.createElement('script'),
						placeTag = document.getElementsByTagName("script")[0];
					scriptTag.type = 'text/javascript';
					scriptTag.async = true;
					scriptTag.src = tagSrc;
					s.parentNode.insertBefore(scriptTag, s);
				}
	
				function loadGPT() {
					if (!window.advBidxc.isAdServerLoaded) {
						loadScript('//www.googletagservices.com/tag/js/gpt.js');
						window.advBidxc.isAdServerLoaded = true;
					}
				}
	
				window.advBidxc = window.advBidxc || {};
				window.advBidxc.renderAd = function () {};
				window.advBidxc.startTime = new Date().getTime();
				window.advBidxc.customerId = "8CU2553YN"; //Media.net Customer ID
				window.advBidxc.timeout = 300;
				window.advBidxc.loadGPT = setTimeout(loadGPT, window.advBidxc.timeout);
	
				// append index
				var a = document.createElement("script");
				a.type = "text/javascript";
				a.async = true;
				a.src = indexBid;
				s.parentNode.insertBefore(a, s);
	
				// append media.net
				var b = document.createElement("script");
				b.type = "text/javascript";
				b.async = true;
				b.src = mediaDotNet;
				s.parentNode.insertBefore(b, s);
			})();
		</script>
	
		<!-- Amazon Bidder version 2 Initialization Script -->
		<script type="text/javascript">
			//Load the ApsTag JavaScript Library
			!function(a9,a,p,s,t,A,g){if(a[a9])return;function q(c,r){a[a9]._Q.push([c,r])}a[a9]={init:function(){q("i",arguments)},fetchBids:function(){q("f",arguments)},setDisplayBids:function(){},targetingKeys:function(){return[]},_Q:[]};A=p.createElement(s);A.async=!0;A.src=t;g=p.getElementsByTagName(s)[0];g.parentNode.insertBefore(A,g)}("apstag",window,document,"script","//c.amazon-adsystem.com/aax2/apstag.js");
	
			apstag.init({
				pubID: '3030',
				adServer: 'googletag'
			});
		</script>
	
	
		<script type="text/json" id="trending-link-json">
		</script>
	
	<!--esi
	<script id="user-info-data" type="application/json">
	<esi:include src="/svc/web-products/userinfo-v3.json" />
	</script>
	-->
	<script id="magnum-feature-flags" type="application/json">["limitFabrikSave","moreFollowSuggestions","unfollowComments","scoopTool","videoVHSCover","videoVHSShareTools","videoVHSLive","videoVHSNewControls","videoVHSEmbeddedOnly","allTheEmphases","freeTrial","dedupeWhatsNext","trendingPageLinks","sprinklePaidPost","headerBidder","standardizeWhatsNextCollection","onlyLayoutA","simple","simpleExtendedByline","collectionsWhatsNext","mobileMediaViewer","podcastInLede","MovieTickets","MoreInSubsectionFix","seriesIssueMarginalia","serverSideCollectionUrls","multipleBylines","fabrikWebsocketsOnly","translationLinks","papihttps","verticalFullBleed","updateRestaurantReservations","mapDining","newsletterInlineModule","PersonalizationApiUpdate","apsTagEnabled","removeInternationalEdition","headlineBalancer","clientSideABRA","newsletterTitleTracking","surveyInterstitial","removeFBMessengerPromo","removeMarginaliaAbTest","paidpostSprinklingFix","abraOverrideVersion","headlineBalancerEverywhere","signupFavor","lazyloadOakImages","mapleFreeTrial","adQuerySupportForMultipleUserAddedKeywords","removeSectionDependentAdLogicForJanuary","oakUpdateAdStride","autoPlaceNewsletter","piiBlockDFP","adExclusiveIERibbonCheck","FlexAdDoubleHeightTracking","didScroll","supportedByAdFullBleed","anchorsHttps","oakBylineMargin","TopFlexAdSiteWide","BundlePayFlow","loadArticleOptimizely","oakStyleTouchups","oakBasicHeader","serveOakOnVI","opinionHeadline","indexAsHeaderBidder","newsletterRegiTest","bookReviewBuyButton","caslOpt","sectionBasedAmp","abraInCustomParams","jkiddScript","blueKai","showNewDealbookLogo","commentsUserDataUpdateDisabled","tMagazineFontTest","disableServeAsNyt4","signupModuleTest"]</script>
	<script>
	var require = {
		baseUrl: 'https://g1.nyt.com/assets/',
		waitSeconds: 20,
		paths: {
			'foundation': 'article/20180322-142545/js/foundation',
			'shared': 'article/20180322-142545/js/shared',
			'article': 'article/20180322-142545/js/article',
			'application': 'article/20180322-142545/js/article/story',
			'videoFactory': 'https://static01.nyt.com/js2/build/video/2.0/videofactoryrequire',
			'videoPlaylist': 'https://static01.nyt.com/js2/build/video/players/extended/2.0/appRequire',
			'auth/mtr': 'https://static01.nyt.com/js/mtr',
			'auth/growl': 'https://static01.nyt.com/js/auth/growl/default',
			'vhs': 'https://static01.nyt.com/video/vhs/build/vhs-2.x.min',
			'vhs3': 'https://static01.nyt.com/video-static/vhs3/vhs.min'
		},
		map: {
			'*': {
				'story/main': 'article/story/main'
			}
		}
	};
	</script>
	<!--[if (gte IE 9)|!(IE)]> <!-->
	<script data-main="foundation/main" src="https://g1.nyt.com/assets/article/20180322-142545/js/foundation/lib/framework.js"></script>
	<!--<![endif]-->
	<!--[if lt IE 9]>
	<script>
		require.map['*']['foundation/main'] = 'foundation/legacy_main';
	</script>
	<script data-main="foundation/legacy_main" src="https://g1.nyt.com/assets/article/20180322-142545/js/foundation/lib/framework.js"></script>
	<![endif]-->
		<script>
		require(['foundation/main'], function () {
			require(['auth/mtr', 'auth/growl']);
		});
		</script>
	</head>
	
	<body>		 
	</body>
	</html>

`
)

func TestGetOpenGraphFromUrl(t *testing.T) {

	mockHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockHtml))
	}

	mockServer := httptest.NewServer(http.HandlerFunc(mockHandler))
	defer mockServer.Close()

	res, err := og.GetOpenGraphFromUrl(mockServer.URL)
	if err != nil {
		t.Error("Expected not Err but got Err: ", err)
	}

	if res == nil {
		t.Error("Result expected not nil but got nil")
	}
	if res.Title == "" {
		t.Errorf("Title expected not empty, but got an empty title")
	}
}

func TestOpenGraphFromHtmlString(t *testing.T) {
	res, err := og.GetOpenGraphFromHtml(mockHtml)
	if err != nil {
		t.Error("Expected not Err but got Err: ", err)
	}

	if res == nil {
		t.Error("Result expected not nil but got nil")
	}
	if res.Title == "" {
		t.Errorf("Title expected not empty, but got an empty title")
	}
}
