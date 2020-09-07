package albums

import (
	"bytes"
	"github.com/a-castellano/music-manager-metal-archives-wrapper/types"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetTracks(t *testing.T) {

	albumData := SearchAlbumData{Name: "Soma", URL: "https://www.metal-archives.com/albums/B%C3%B6lzer/Soma/447710", ID: 447710, Year: 2014, Artist: "Bölzer", ArtistID: 3540351548, ArtistURL: "https://www.metal-archives.com/bands/B%C3%B6lzer/3540351548", Type: types.EP}

	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<title>Bölzer - Soma - Encyclopaedia Metallum: The Metal Archives</title>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<link rel="stylesheet" type="text/css" href="/min/index.php?g=css" />
<script type="609698a8fdb6bad6d2a2627d-text/javascript" src="/min/index.php?g=js"></script>
<script type="609698a8fdb6bad6d2a2627d-text/javascript">
			var URL_SITE	= 'https://www.metal-archives.com/';
			var URL_IMAGES	= 'https://www.metal-archives.com/images/';
			var URL_CSS		= 'https://www.metal-archives.com/css/';
			var csrfToken	= '';

			$(document).ready(function() {
				toggleFixedPositioning(1280, 600);
				executeOnAllPages();

											});


		</script>
<noscript>
			<style type="text/css">
				.no-js { display: block !important; }
			</style>
		</noscript>
<!--[if lt IE 8]>
			<script src="http://ie7-js.googlecode.com/svn/version/2.1(beta4)/IE8.js"></script>
			 <link rel="stylesheet" type="text/css" href="https://www.metal-archives.com/css/default/oldie.css" />
		<![endif]-->
<script type="609698a8fdb6bad6d2a2627d-text/javascript">
			  var _gaq = _gaq || [];
			  _gaq.push(['_setAccount', 'UA-4046749-4']);
			  _gaq.push(['_trackPageview']);

			  (function() {
			    var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
			    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
			    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
			  })();
		</script>
</head>
<body>
<div id="wrapper">
<div id="header">
<a href="https://www.metal-archives.com/" id="MA_logo">Metal Archives</a>
<div class="clear loading"><img src="https://www.metal-archives.com/images/loading.gif" alt="loading" width="12" height="12" /> loading...</div>
<div id="search_box">
<form id="search_form" action="https://www.metal-archives.com/search" method="get">
<div>
<label>Search:</label><br />
<input name="searchString" type="text" id="searchQueryBox" size="25" value="" tabindex="1" />
<select name="type" tabindex="2">
<option value="band_name">Band name</option>
<option value="band_genre">Music genre</option>
<option value="band_themes">Lyrical themes</option>
<option value="album_title">Album title</option>
<option value="song_title">Song title</option>
<option value="label_name">Label</option>
<option value="artist_alias">Artist</option>
<option value="user_name">User profile</option>
<option value="google">Google</option>
</select>
<br />
<a href="https://www.metal-archives.com/search/advanced/?searchString=&type=" class="float_left" tabindex="4">Advanced search</a>
<button type="submit" class="btn_submit float_right" tabindex="3">Submit</button>
</div>
</form>
</div>
<div id="top_menu_box">
<ul class="menu_style_1">
<li><a href="https://www.metal-archives.com/content/help">Help</a></li>
<li><a href="https://www.metal-archives.com/content/rules">Rules</a></li>
<li><a href="https://www.metal-archives.com/board/" target="_top">Forum</a></li>
</ul>
<ul class="menu_style_2">
<li><a href="https://www.metal-archives.com/content/faq">FAQ</a></li>
<li><a href="https://www.metal-archives.com/content/support">Support Us</a></li>
<li><a href="https://www.metal-archives.com/content/tools">Add-ons</a></li>
</ul>
</div>
</div>
<div id="left_col">
<div id="member_box">
<form name="login_form" id="login_form" action="https://www.metal-archives.com/authentication/login" method="post">
<div>
<label>Username</label><br />
<input type="text" name="loginUsername" />
<label>Password</label><br />
<input type="password" name="loginPassword" />
<input type="hidden" name="origin" value="/albums/B%C3%B6lzer/Soma/447710" />
<button type="submit" class="btn_login float_right">Login</button>
</div>
</form>
<a href="https://www.metal-archives.com/user/signup" class="float-left writeAction">Register</a>
<span class="float_right"><a href="https://www.metal-archives.com/user/forgot-password" class="float-left writeAction">Forgot login?</a> </span>
<div class="clear"> </div>
</div>
<div id="left_menu_box">
<ul class="menu_style_3 block_spacer_20">
<li>
Bands
<ul>
<li><a href="https://www.metal-archives.com/browse/letter">alphabetical</a></li>
<li><a href="https://www.metal-archives.com/browse/country">country</a></li>
<li><a href="https://www.metal-archives.com/browse/genre">genre</a></li>
</ul>
</li>
<li>
Labels
<ul>
<li><a href="https://www.metal-archives.com/label">alphabetical</a></li>
<li><a href="https://www.metal-archives.com/label/country">country</a></li>
</ul>
</li>
<li><a href="https://www.metal-archives.com/review/browse">Reviews</a></li>
</ul>
<ul class="menu_style_3 block_spacer_20">
<li><a href="https://www.metal-archives.com/artist/rip">R.I.P.</a></li>
<li><a href="https://www.metal-archives.com/band/random">Random Band</a></li>
<li><a href="https://www.metal-archives.com/user/list">User rankings</a></li>
<li><a href="https://www.metal-archives.com/news">News archive</a></li>
</ul>
<ul class="menu_style_3 ornement_30">
<li><a href="https://www.metal-archives.com/report/list">Reports</a></li>
<li><a href="https://www.metal-archives.com/todo">Contribute / To do</a></li>
</ul>
</div>
<div id="left_text_box">
<p class="small_text center"> &copy; 2002-2020<br />Encyclopaedia Metallum </p>
<p class="small_text center">Best viewed<br /> <a href="http://www.getfirefox.com" target="_blank">without</a> Internet Explorer, <br />in 1280 x 960 resolution<br /> or higher.</p>
</div>
<div>
<p class="small_text center"><a href="https://www.metal-archives.com/content/pp">Privacy Policy</a></p>
</div>
</div>
<div id="content_wrapper">
<script type="609698a8fdb6bad6d2a2627d-text/javascript" src="https://www.metal-archives.com/js/jquery/jquery.form.js"></script>
<script type="609698a8fdb6bad6d2a2627d-text/javascript">

	$(document).ready(function() {
		$(".no-js").removeClass("no-js");
		$('#album_tabs').tabs({ beforeLoad: cacheTab }).show();
		$('#album_members').tabs();
		$('.albumMenu').click(toggleAlbumMenu);
		 $('#album_members').tabs( { active : 1 } );							toggleLyrics(0);
			});

	function hideToolMenu(event) {
		 $('.toolMenu-content').hide();
	}

	function makeChildVersion() {
		var releaseId = 447710;
		var newParentId = prompt("Make a child of which album ID?");

		if(newParentId != null) {

			$('.loading').show();
			$.post(URL_SITE + 'review/ajax-validate-move/json/1', { 'releaseId' : newParentId, "oldReleaseId": releaseId }, function(result) {
				$('.loading').hide();
				if(result.reviewConflict) {
					alert("This action would cause at least two reviews by the same user to be listed for one album. This is not allowed. One of the reviews will need to be deleted first.");
				}
				else if(result.success) {
					var confirmMove = confirm("Make this album (and all its own versions, if any) a version of " + result.band + " - " + result.release + "?");

					if(confirmMove) {
						var params = { 'releaseId' : releaseId, "newParentId" : newParentId};
						$('.loading').show();
						$.post(URL_SITE + 'release/ajax-make-child/json/1', params, function(result) {
							$('.loading').hide();
							if(result.success) {
								location.href = "https://www.metal-archives.com/albums/B%C3%B6lzer/Soma/447710";
							}
							else {
								alert("An error occurred: " + result.error);
							}
						}, 'json');
					}
				}
				else {
					alert("An error occurred: " + result.error);
				}
			}, 'json');
		}
	}


</script>
<div id="album_sidebar">
<a href="https://www.metal-archives.com/bands/B%C3%B6lzer/3540351548#band_tab_discography">Bölzer</a>
> Soma

<div class="album_img">
<a class="image" id="cover" title="Bölzer - Soma" href="https://www.metal-archives.com/images/4/4/7/7/447710.jpg?5809"><img src="https://www.metal-archives.com/images/4/4/7/7/447710.jpg?5809" title="Click to zoom" alt="Bölzer - Soma" border="0" /></a>
</div>

<div id="affiliation-links">
<span class="title">Buy from...</span>
<br />
<ul id="buyLinks" class="menu_style_5">
<li>
<a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=29&rid=447710">eBay</a>
more... <a href="#more" onclick="if (!window.__cfRLUnblockHandlers) return false; $('#sub29').toggle(); return false;" data-cf-modified-609698a8fdb6bad6d2a2627d-="">>></a><br />
<ul id="sub29" style="display: none; margin-top: 0px; margin-bottom: 0px;">
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=23&rid=447710">eBay Canada</a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=25&rid=447710">eBay UK</a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=58&rid=447710">eBay Germany</a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=24&rid=447710">eBay France</a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=37&rid=447710">eBay Spain</a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=35&rid=447710">eBay Belgium</a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=26&rid=447710">eBay Netherlands</a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=27&rid=447710">eBay Italy</a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=28&rid=447710">eBay Australia</a></li>
</ul>
</li>
</ul>
</div>

<table cellspacing="1" cellpadding="1" width="262" class="chronology">
<tr class="hidden">
<th class="spacer"></th>
<th colspan="2"></th>
<th class="spacer"></th>
</tr>
<tr>
<th colspan="4"><span>Bölzer</span> discography (misc)</th>
</tr>
<tr class="prevNext">
<td class="arrows"><a href="https://www.metal-archives.com/albums/B%C3%B6lzer/Aura/376088" title="Aura (EP)"><</a> </td>
<td class="prev">
<a href="https://www.metal-archives.com/albums/B%C3%B6lzer/Aura/376088" title="Aura (EP)">Aura</a><br />(2013)
</td>
<td class="next">
<a href="https://www.metal-archives.com/albums/B%C3%B6lzer/C.H.A.O.S./661984" title="C.H.A.O.S. (Split)">C.H.A.O.S.</a><br />(2017)
</td>
<td class="arrows"><a href="https://www.metal-archives.com/albums/B%C3%B6lzer/C.H.A.O.S./661984" title="C.H.A.O.S. (Split)">></a> </td>
</tr>
<tr>
<th colspan="4"><span>Bölzer</span> discography (all)</th>
</tr>
<tr class="prevNext">
<td class="arrows"><a href="https://www.metal-archives.com/albums/B%C3%B6lzer/Aura/376088" title="Aura (EP)"><</a> </td>
<td class="prev">
<a href="https://www.metal-archives.com/albums/B%C3%B6lzer/Aura/376088" title="Aura (EP)">Aura</a><br />(2013)
</td>
<td class="next">
<a href="https://www.metal-archives.com/albums/B%C3%B6lzer/Hero/604091" title="Hero (Full-length)">Hero</a><br />(2016)
</td>
<td class="arrows"><a href="https://www.metal-archives.com/albums/B%C3%B6lzer/Hero/604091" title="Hero (Full-length)">></a> </td>
</tr>
</table>
</div>
<div id="album_content">
<div class="tool_strip right">
<ul>
<li><a title="Report a mistake or additional information for this page" href="javascript:popupReportDialog(4, 447710);" class="btn_report_error writeAction"> </a></li>
<li>
<a href="https://www.metal-archives.com/bands/B%C3%B6lzer/3540351548" title="Back to Bölzer" class="btn_back"> </a>
</li>
</ul>
</div>

<div id="album_info">

<h1 class="album_name"><a href="https://www.metal-archives.com/albums/B%C3%B6lzer/Soma/447710">Soma</a></h1>
<h2 class="band_name">
<a href="https://www.metal-archives.com/bands/B%C3%B6lzer/3540351548">Bölzer</a>
</h2>

<div class="clear block_spacer_5"></div>
<div id="message"> </div>

<div class="clear block_spacer_20"></div>
<dl class="float_left">
<dt>Type:</dt>
<dd>EP</dd>
<dt>Release date:</dt>
<dd>August 11th, 2014</dd>
<dt>Catalog ID:</dt>
<dd>IP058</dd>
</dl>
<dl class="float_right">
<dt>Label:</dt>
<dd><a href="https://www.metal-archives.com/labels/Invictus_Productions/815#label_tabs_albums">Invictus Productions</a></dd>
<dt>Format:</dt>
<dd>CD</dd>
<dt>Reviews:</dt>
<dd>
7 <a href="https://www.metal-archives.com/reviews/B%C3%B6lzer/Soma/447710/">reviews</a> (avg. 83%)
</dd>
</dl>
</div>
<div id="album_tabs" class="clear tabs block_spacer_top_20 ui-tabs no-js">
<ul class="ui-tabs-nav">
<li><a href="#album_tabs_tracklist">Songs</a></li>
<li><a href="#album_tabs_lineup">Lineup</a></li>
<li><a href="https://www.metal-archives.com/release/ajax-versions/current/447710/parent/447710">Other versions</a></li> <li><a href="#album_tabs_reviews">Reviews</a></li> <li><a href="#album_tabs_notes">Additional notes</a></li> </ul>

<div id="album_tabs_tracklist" class="ui-tabs-hide">
<div id="album_songs" class="tabs2lvl">

</div>
<div class="ui-tabs-panel-content block_spacer_top_20">
<script type="609698a8fdb6bad6d2a2627d-text/javascript">
function toggleLyrics(songId) {
    var lyricsRow = $('#song' + songId);
    lyricsRow.toggle();
    var lyrics = $('#lyrics_' + songId);
	if (lyrics.html() == '(loading lyrics...)') {
    	var realId = songId;
		if(!$.isNumeric(songId.substring(songId.length -1, songId.length))) {
			realId = songId.substring(0, songId.length -1);
		}
		lyrics.load(URL_SITE + "release/ajax-view-lyrics/id/" + realId);
    }
    // toggle link
    var linkLabel = "lyrics";
    $("#lyricsButton" + songId).text(lyricsRow.css("display") == "none" ? "Show " + linkLabel : "Hide " + linkLabel);
    return false;
}

</script>
<table class="display table_lyrics" cellpadding="0" cellspacing="0">
<tbody>
<tr class="even">
<td width="20"><a name="3074706" class="anchor"> </a>1.</td>
<td class="wrapWords">
Steppes
</td>
<td align="right">05:34</td>
<td nowrap="nowrap">&nbsp;
<a id="lyricsButton3074706" href="#3074706" onclick="if (!window.__cfRLUnblockHandlers) return false; toggleLyrics('3074706'); return false;" data-cf-modified-609698a8fdb6bad6d2a2627d-="">Show lyrics</a>
</td>
</tr>
<tr id="song3074706" class="displayNone" height="0">
<td>&nbsp;</td>
<td colspan="3" id="lyrics_3074706">(loading lyrics...)</td>
</tr>
<tr class="odd">
<td width="20"><a name="3074707" class="anchor"> </a>2.</td>
<td class="wrapWords">
Labyrinthian Graves
</td>
<td align="right">12:28</td>
<td nowrap="nowrap">&nbsp;
<a id="lyricsButton3074707" href="#3074707" onclick="if (!window.__cfRLUnblockHandlers) return false; toggleLyrics('3074707'); return false;" data-cf-modified-609698a8fdb6bad6d2a2627d-="">Show lyrics</a>
</td>
</tr>
<tr id="song3074707" class="displayNone" height="0">
<td>&nbsp;</td>
<td colspan="3" id="lyrics_3074707">(loading lyrics...)</td>
</tr>
<tr>
<td colspan="2">&nbsp;</td>
<td align="right"><strong>18:02</strong></td>
<td>&nbsp;</td>
</tr>
</tbody>
</table>
</div>
</div>


<div id="album_tabs_lineup">
<div id="album_members" class="tabs2lvl">
<ul>
<li><a href="#album_all_members_lineup">Complete lineup</a></li> <li><a href="#album_members_lineup">Band members</a></li> <li><a href="#album_members_misc">Other staff</a></li> </ul>

<div id="album_all_members_lineup">
<div class="ui-tabs-panel-content">
<table class="display lineupTable" cellpadding="0" cellspacing="0">
<tr class="lineupHeaders">
<td colspan="2" align="right">
Band members
</td>
</tr>
<tr class="lineupRow">
<td width="300" valign="top">
<a href="https://www.metal-archives.com/artists/HzR/450000">HzR</a>
</td>
<td>
Drums </td>
</tr>
<tr class="lineupRow">
<td width="300" valign="top">
<a href="https://www.metal-archives.com/artists/Okoi_Thierry_Jones/260204">KzR</a>
</td>
<td>
Guitars, Vocals </td>
</tr>
<tr class="lineupHeaders">
<td colspan="2" align="right">
Miscellaneous staff
</td>
</tr>
<tr class="lineupRow">
<td width="300" valign="top">
<a href="https://www.metal-archives.com/artists/Alexander_L._Brown/46749">Alexander L. Brown</a>
</td>
<td>
Artwork </td>
</tr>
<tr class="lineupRow">
<td width="300" valign="top">
<a href="https://www.metal-archives.com/artists/C._Sinclair/25769">Cam Sinclair</a>
</td>
<td>
Mastering </td>
</tr>
</table>
</div>
</div>

<div id="album_members_lineup">
<div class="ui-tabs-panel-content">
<table class="display lineupTable" cellpadding="0" cellspacing="0">
<tr class="lineupRow">
<td width="300" valign="top">
<a href="https://www.metal-archives.com/artists/HzR/450000">HzR</a>
</td>
<td>
Drums </td>
</tr>
<tr class="lineupRow">
<td width="300" valign="top">
<a href="https://www.metal-archives.com/artists/Okoi_Thierry_Jones/260204">KzR</a>
 </td>
<td>
Guitars, Vocals </td>
</tr>
</table>
</div>
</div>

<div id="album_members_misc">
<div class="ui-tabs-panel-content">
<table class="display lineupTable" cellpadding="0" cellspacing="0">
<tr class="lineupRow">
<td width="300" valign="top">
<a href="https://www.metal-archives.com/artists/Alexander_L._Brown/46749">Alexander L. Brown</a>
</td>
<td>
Artwork </td>
</tr>
<tr class="lineupRow">
<td width="300" valign="top">
<a href="https://www.metal-archives.com/artists/C._Sinclair/25769">Cam Sinclair</a>
</td>
<td>
Mastering </td>
</tr>
</table>
</div>
</div>
</div>
</div>





<div id="album_tabs_reviews" class="ui-tabs-hide">
<div id="album_reviews" class="tabs2lvl">
<div class="tool_strip top right writeAction">
<ul>
<li><a href="https://www.metal-archives.com/review/write/releaseId/447710" class="btn_add">Add</a></li>
</ul>
</div>
</div>
<div class="ui-tabs-panel-content block_spacer_top_36">
<table id="review_list" class="display" cellpadding="0" cellspacing="0">
<tr class="even">
<td nowrap="nowrap"><a href="https://www.metal-archives.com/reviews/B%C3%B6lzer/Soma/447710/DSOfan97/362569" title="Read" class="iconContainer ui-state-default ui-corner-all"><span class="ui-icon ui-icon-search">Read</span></a></td>
<td>The Swiss bulldozer is back!</td>
<td>88%</td>
<td><a href="https://www.metal-archives.com/users/DSOfan97" class="profileMenu">DSOfan97</a></td>
<td>June 6th, 2015</td>
</tr>
<tr class="odd">
<td nowrap="nowrap"><a href="https://www.metal-archives.com/reviews/B%C3%B6lzer/Soma/447710/PassiveMetalhead/342911" title="Read" class="iconContainer ui-state-default ui-corner-all"><span class="ui-icon ui-icon-search">Read</span></a></td>
<td>A Force To Be Reckoned With</td>
<td>88%</td>
<td><a href="https://www.metal-archives.com/users/PassiveMetalhead" class="profileMenu">PassiveMetalhead</a></td>
<td>May 18th, 2015</td>
</tr>
<tr class="even">
<td nowrap="nowrap"><a href="https://www.metal-archives.com/reviews/B%C3%B6lzer/Soma/447710/dismember_marcin/217143" title="Read" class="iconContainer ui-state-default ui-corner-all"><span class="ui-icon ui-icon-search">Read</span></a></td>
<td>Bolzer - Soma</td>
<td>90%</td>
<td><a href="https://www.metal-archives.com/users/dismember_marcin" class="profileMenu">dismember_marcin</a></td>
<td>January 14th, 2015</td>
</tr>
<tr class="odd">
<td nowrap="nowrap"><a href="https://www.metal-archives.com/reviews/B%C3%B6lzer/Soma/447710/Achintya_Venkatesh/316729" title="Read" class="iconContainer ui-state-default ui-corner-all"><span class="ui-icon ui-icon-search">Read</span></a></td>
<td>Sonic soaring</td>
<td>83%</td>
<td><a href="https://www.metal-archives.com/users/Achintya%20Venkatesh" class="profileMenu">Achintya Venkatesh</a></td>
<td>September 26th, 2014</td>
 </tr>
<tr class="even">
<td nowrap="nowrap"><a href="https://www.metal-archives.com/reviews/B%C3%B6lzer/Soma/447710/Witchfvcker/335640" title="Read" class="iconContainer ui-state-default ui-corner-all"><span class="ui-icon ui-icon-search">Read</span></a></td>
<td>The wrath of an angry God</td>
<td>85%</td>
<td><a href="https://www.metal-archives.com/users/Witchfvcker" class="profileMenu">Witchfvcker</a></td>
<td>August 30th, 2014</td>
</tr>
<tr class="odd">
<td nowrap="nowrap"><a href="https://www.metal-archives.com/reviews/B%C3%B6lzer/Soma/447710/ThrashManiacAYD/205006" title="Read" class="iconContainer ui-state-default ui-corner-all"><span class="ui-icon ui-icon-search">Read</span></a></td>
<td>Bölzer - Soma</td>
<td>75%</td>
<td><a href="https://www.metal-archives.com/users/ThrashManiacAYD" class="profileMenu">ThrashManiacAYD</a></td>
<td>August 10th, 2014</td>
</tr>
<tr class="even">
<td nowrap="nowrap"><a href="https://www.metal-archives.com/reviews/B%C3%B6lzer/Soma/447710/autothrall/192699" title="Read" class="iconContainer ui-state-default ui-corner-all"><span class="ui-icon ui-icon-search">Read</span></a></td>
<td>These two things are not equal</td>
<td>72%</td>
<td><a href="https://www.metal-archives.com/users/autothrall" class="profileMenu">autothrall</a></td>
<td>August 5th, 2014</td>
</tr>
</table>
</div>
</div>


<div id="album_tabs_notes" class="ui-tabs-hide">
<div id="album_notes" class="tabs2lvl"></div>
<div class="ui-tabs-panel-content  block_spacer_top_20">
<p class="block_spacer_20"></p>
<p class="title_comment">Recording information:</p>
<p class="block_spacer_20">Captured at Osa Crypt, Turicum, Summer MMXIII.<br />
Anointed at Temple of Sol.</p>
</div>
</div>

</div>

<div id="auditTrail">
<table>
<tr>
<td>Added by: <a href="https://www.metal-archives.com/users/Amoebic_D" class="profileMenu">Amoebic_D</a></td>
<td align="right">Modified by: <a href="https://www.metal-archives.com/users/Sang%20Dalang%20Abu" class="profileMenu">Sang Dalang Abu</a></td>
</tr>
<tr>
<td>Added on: 2014-08-21 12:58:09</td>
<td align="right">Last modified on: 2018-07-05 11:53:49</td>
</tr>
<tr>
<td valign="top">
&nbsp;
</td>
<td align="right" valign="top">
</td>
</tr>
</table>
</div>

</div>
</div>
</div>
<script src="https://ajax.cloudflare.com/cdn-cgi/scripts/7089c43e/cloudflare-static/rocket-loader.min.js" data-cf-settings="609698a8fdb6bad6d2a2627d-|49" defer=""></script></body>
</html>
	`))}}}

	tracks, cover, err := GetAlbumInfo(client, albumData)

	if err != nil {
		t.Errorf("SearchAlbumData shouldn't fail.")
	}

	if len(tracks) != 2 {
		t.Errorf("Soma by Bölzer has only 2 tracks not %d.", len(tracks))
	}

	if tracks[0].Name != "Steppes" {
		t.Errorf("Soma by Bölzer first track is called 'Steppes', not '%s'.", tracks[0].Name)
	}

	if tracks[0].Hours != 0 {
		t.Errorf("Steppes by Bölzer hours should be 0, not '%d'.", tracks[0].Hours)
	}

	if tracks[0].Minutes != 5 {
		t.Errorf("Steppes by Bölzer minutes should be 5, not '%d'.", tracks[0].Minutes)
	}

	if tracks[0].Seconds != 34 {
		t.Errorf("Steppes by Bölzer seconds should be 34, not '%d'.", tracks[0].Seconds)
	}

	if tracks[1].Name != "Labyrinthian Graves" {
		t.Errorf("Soma by Bölzer second track is called 'Labyrinthian Graves', not '%s'.", tracks[1].Name)
	}

	if tracks[1].Hours != 0 {
		t.Errorf("Labyrinthian Graves by Bölzer hours should be 0, not '%d'.", tracks[1].Hours)
	}

	if tracks[1].Minutes != 12 {
		t.Errorf("Labyrinthian Graves by Bölzer minutes should be 12, not '%d'.", tracks[1].Minutes)
	}

	if tracks[1].Seconds != 28 {
		t.Errorf("Labyrinthian Graves by Bölzer seconds should be 28, not '%d'.", tracks[1].Seconds)
	}

	if cover != "https://www.metal-archives.com/images/4/4/7/7/447710.jpg" {
		t.Errorf("Soma by Bölzer has cover located in 'https://www.metal-archives.com/images/4/4/7/7/447710.jpg', not %s'.", cover)
	}

}

func TestLongerTracks(t *testing.T) {

	albumData := SearchAlbumData{Name: "The Hunt", URL: "https://www.metal-archives.com/albums/Fauna/The_Hunt/189275", ID: 189275, Year: 2007, Artist: "Fauna", ArtistID: 121144, ArtistURL: "https://www.metal-archives.com/bands/Fauna/121144", Type: types.FullLength}

	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<title>Fauna - The Hunt - Encyclopaedia Metallum: The Metal Archives</title>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<link rel="stylesheet" type="text/css" href="/min/index.php?g=css" />
<script type="fa3d21b0fdc73efe12102b0f-text/javascript" src="/min/index.php?g=js"></script>
<script type="fa3d21b0fdc73efe12102b0f-text/javascript">
			var URL_SITE	= 'https://www.metal-archives.com/';
			var URL_IMAGES	= 'https://www.metal-archives.com/images/';
			var URL_CSS		= 'https://www.metal-archives.com/css/';
			var csrfToken	= '';

			$(document).ready(function() {
				toggleFixedPositioning(1280, 600);
				executeOnAllPages();

											});


		</script>
<noscript>
			<style type="text/css">
				.no-js { display: block !important; }
			</style>
		</noscript>
<!--[if lt IE 8]>
			<script src="http://ie7-js.googlecode.com/svn/version/2.1(beta4)/IE8.js"></script>
			 <link rel="stylesheet" type="text/css" href="https://www.metal-archives.com/css/default/oldie.css" />
		<![endif]-->
<script type="fa3d21b0fdc73efe12102b0f-text/javascript">
			  var _gaq = _gaq || [];
			  _gaq.push(['_setAccount', 'UA-4046749-4']);
			  _gaq.push(['_trackPageview']);

			  (function() {
			    var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
			    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
			    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
			  })();
		</script>
</head>
<body>
<div id="wrapper">
<div id="header">
<a href="https://www.metal-archives.com/" id="MA_logo">Metal Archives</a>
<div class="clear loading"><img src="https://www.metal-archives.com/images/loading.gif" alt="loading" width="12" height="12" /> loading...</div>
<div id="search_box">
<form id="search_form" action="https://www.metal-archives.com/search" method="get">
<div>
<label>Search:</label><br />
<input name="searchString" type="text" id="searchQueryBox" size="25" value="" tabindex="1" />
<select name="type" tabindex="2">
<option value="band_name" selected="selected">Band name</option>
<option value="band_genre">Music genre</option>
<option value="band_themes">Lyrical themes</option>
<option value="album_title">Album title</option>
<option value="song_title">Song title</option>
<option value="label_name">Label</option>
<option value="artist_alias">Artist</option>
<option value="user_name">User profile</option>
<option value="google">Google</option>
</select>
<br />
<a href="https://www.metal-archives.com/search/advanced/?searchString=&type=band_name" class="float_left" tabindex="4">Advanced search</a>
<button type="submit" class="btn_submit float_right" tabindex="3">Submit</button>
</div>
</form>
</div>
<div id="top_menu_box">
<ul class="menu_style_1">
<li><a href="https://www.metal-archives.com/content/help">Help</a></li>
<li><a href="https://www.metal-archives.com/content/rules">Rules</a></li>
<li><a href="https://www.metal-archives.com/board/" target="_top">Forum</a></li>
</ul>
<ul class="menu_style_2">
<li><a href="https://www.metal-archives.com/content/faq">FAQ</a></li>
<li><a href="https://www.metal-archives.com/content/support">Support Us</a></li>
<li><a href="https://www.metal-archives.com/content/tools">Add-ons</a></li>
</ul>
</div>
</div>
<div id="left_col">
<div id="member_box">
<form name="login_form" id="login_form" action="https://www.metal-archives.com/authentication/login" method="post">
<div>
<label>Username</label><br />
<input type="text" name="loginUsername" />
<label>Password</label><br />
<input type="password" name="loginPassword" />
<input type="hidden" name="origin" value="/albums/Fauna/The_Hunt/189275" />
<button type="submit" class="btn_login float_right">Login</button>
</div>
</form>
<a href="https://www.metal-archives.com/user/signup" class="float-left writeAction">Register</a>
<span class="float_right"><a href="https://www.metal-archives.com/user/forgot-password" class="float-left writeAction">Forgot login?</a> </span>
<div class="clear"> </div>
</div>
<div id="left_menu_box">
<ul class="menu_style_3 block_spacer_20">
<li>
Bands
<ul>
<li><a href="https://www.metal-archives.com/browse/letter">alphabetical</a></li>
<li><a href="https://www.metal-archives.com/browse/country">country</a></li>
<li><a href="https://www.metal-archives.com/browse/genre">genre</a></li>
</ul>
</li>
<li>
Labels
<ul>
<li><a href="https://www.metal-archives.com/label">alphabetical</a></li>
<li><a href="https://www.metal-archives.com/label/country">country</a></li>
</ul>
</li>
<li><a href="https://www.metal-archives.com/review/browse">Reviews</a></li>
</ul>
<ul class="menu_style_3 block_spacer_20">
<li><a href="https://www.metal-archives.com/artist/rip">R.I.P.</a></li>
<li><a href="https://www.metal-archives.com/band/random">Random Band</a></li>
<li><a href="https://www.metal-archives.com/user/list">User rankings</a></li>
<li><a href="https://www.metal-archives.com/news">News archive</a></li>
</ul>
<ul class="menu_style_3 ornement_30">
<li><a href="https://www.metal-archives.com/report/list">Reports</a></li>
<li><a href="https://www.metal-archives.com/todo">Contribute / To do</a></li>
</ul>
</div>
<div id="left_text_box">
<p class="small_text center"> &copy; 2002-2020<br />Encyclopaedia Metallum </p>
<p class="small_text center">Best viewed<br /> <a href="http://www.getfirefox.com" target="_blank">without</a> Internet Explorer, <br />in 1280 x 960 resolution<br /> or higher.</p>
</div>
<div>
<p class="small_text center"><a href="https://www.metal-archives.com/content/pp">Privacy Policy</a></p>
</div>
</div>
<div id="content_wrapper">
<script type="fa3d21b0fdc73efe12102b0f-text/javascript" src="https://www.metal-archives.com/js/jquery/jquery.form.js"></script>
<script type="fa3d21b0fdc73efe12102b0f-text/javascript">

	$(document).ready(function() {
		$(".no-js").removeClass("no-js");
		$('#album_tabs').tabs({ beforeLoad: cacheTab }).show();
		$('#album_members').tabs();
		$('.albumMenu').click(toggleAlbumMenu);
									toggleLyrics(0);
			});

	function hideToolMenu(event) {
		 $('.toolMenu-content').hide();
	}

	function makeChildVersion() {
		var releaseId = 189275;
		var newParentId = prompt("Make a child of which album ID?");

		if(newParentId != null) {

			$('.loading').show();
			$.post(URL_SITE + 'review/ajax-validate-move/json/1', { 'releaseId' : newParentId, "oldReleaseId": releaseId }, function(result) {
				$('.loading').hide();
				if(result.reviewConflict) {
					alert("This action would cause at least two reviews by the same user to be listed for one album. This is not allowed. One of the reviews will need to be deleted first.");
				}
				else if(result.success) {
					var confirmMove = confirm("Make this album (and all its own versions, if any) a version of " + result.band + " - " + result.release + "?");

					if(confirmMove) {
						var params = { 'releaseId' : releaseId, "newParentId" : newParentId};
						$('.loading').show();
						$.post(URL_SITE + 'release/ajax-make-child/json/1', params, function(result) {
							$('.loading').hide();
							if(result.success) {
								location.href = "https://www.metal-archives.com/albums/Fauna/The_Hunt/189275";
							}
							else {
								alert("An error occurred: " + result.error);
							}
						}, 'json');
					}
				}
				else {
					alert("An error occurred: " + result.error);
				}
			}, 'json');
		}
	}


</script>
<div id="album_sidebar">
<a href="https://www.metal-archives.com/bands/Fauna/121144#band_tab_discography">Fauna</a>
> The Hunt

<div class="album_img">
<a class="image" id="cover" title="Fauna - The Hunt" href="https://www.metal-archives.com/images/1/8/9/2/189275.jpg"><img src="https://www.metal-archives.com/images/1/8/9/2/189275.jpg" title="Click to zoom" alt="Fauna - The Hunt" border="0" /></a>
</div>

<div id="affiliation-links">
<span class="title">Buy from...</span>
<br />
<ul id="buyLinks" class="menu_style_5">
<li>
<a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=29&rid=189275">eBay</a>
more... <a href="#more" onclick="if (!window.__cfRLUnblockHandlers) return false; $('#sub29').toggle(); return false;" data-cf-modified-fa3d21b0fdc73efe12102b0f-="">>></a><br />
<ul id="sub29" style="display: none; margin-top: 0px; margin-bottom: 0px;">
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=23&rid=189275">eBay Canada</a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=25&rid=189275">eBay UK</a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=58&rid=189275">eBay Germany</a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=24&rid=189275">eBay France</a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=37&rid=189275">eBay Spain</a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=35&rid=189275">eBay Belgium</a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=26&rid=189275">eBay Netherlands</a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=27&rid=189275">eBay Italy</a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=28&rid=189275">eBay Australia</a></li>
</ul>
</li>
<li>
<a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=8&id=383234&rid=189275">Amazon.com <img src="https://www.assoc-amazon.com/e/ir?t=encymetatheme-20&l=as2&o=1&a={ID}" width="1" height="1" border="0" alt="" style="border:none !important; margin:0px !important;" /></a>
more... <a href="#more" onclick="if (!window.__cfRLUnblockHandlers) return false; $('#sub8').toggle(); return false;" data-cf-modified-fa3d21b0fdc73efe12102b0f-="">>></a><br />
<ul id="sub8" style="display: none; margin-top: 0px; margin-bottom: 0px;">
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=9&id=383234&rid=189275">Amazon UK <img src="https://www.assoc-amazon.co.uk/e/ir?t=encyclometa0e-21&l=as2&o=2&a={ID}" width="1" height="1" border="0" alt="" style="border:none !important; margin:0px !important;" /></a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=10&id=383234&rid=189275">Amazon Canada <img src="https://www.assoc-amazon.ca/e/ir?t=encyclometall-20&l=as2&o=15&a={ID}" width="1" height="1" border="0" alt="" style="border:none !important; margin:0px !important;" /></a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=11&id=383234&rid=189275">Amazon France <img style="text-decoration:none;border:0;padding:0;margin:0;" src="https://rover.ebay.com/roverimp/1/709-53476-19255-0/1?ff3=9&pub=5574632126&toolid=10001&campid=5335819532&customid={BAND}&uq={BAND}&mpt=[CACHEBUSTER]"></a></li>
<li><a target="_blank" href="https://www.metal-archives.com/affiliate/?pid=12&id=383234&rid=189275">Amazon Germany <img src="https://www.assoc-amazon.de/e/ir?t=encyclometa0a-21&l=as2&o=3&a={ID}" width="1" height="1" border="0" alt="" style="border:none !important; margin:0px !important;" /></a></li>
</ul>
</li>
</ul>
</div>

<table cellspacing="1" cellpadding="1" width="262" class="chronology">
<tr class="hidden">
<th class="spacer"></th>
<th colspan="2"></th>
<th class="spacer"></th>
</tr>
<tr>
<th colspan="4"><span>Fauna</span> discography (main)</th>
</tr>
<tr class="prevNext">
<td class="arrows"><a href="https://www.metal-archives.com/albums/Fauna/Rain/189274" title="Rain (Full-length)"><</a> </td>
<td class="prev">
<a href="https://www.metal-archives.com/albums/Fauna/Rain/189274" title="Rain (Full-length)">Rain</a><br />(2006)
</td>
<td class="next">
<a href="https://www.metal-archives.com/albums/Fauna/Avifauna/361105" title="Avifauna (Full-length)">Avifauna</a><br />(2012)
</td>
<td class="arrows"><a href="https://www.metal-archives.com/albums/Fauna/Avifauna/361105" title="Avifauna (Full-length)">></a> </td>
</tr>
<tr>
<th colspan="4"><span>Fauna</span> discography (all)</th>
</tr>
<tr class="prevNext">
<td class="arrows"><a href="https://www.metal-archives.com/albums/Fauna/Rain/189274" title="Rain (Full-length)"><</a> </td>
<td class="prev">
<a href="https://www.metal-archives.com/albums/Fauna/Rain/189274" title="Rain (Full-length)">Rain</a><br />(2006)
</td>
<td class="next">
<a href="https://www.metal-archives.com/albums/Fauna/Avifauna/361105" title="Avifauna (Full-length)">Avifauna</a><br />(2012)
</td>
<td class="arrows"><a href="https://www.metal-archives.com/albums/Fauna/Avifauna/361105" title="Avifauna (Full-length)">></a> </td>
</tr>
</table>
</div>
<div id="album_content">
<div class="tool_strip right">
<ul>
<li><a title="Report a mistake or additional information for this page" href="javascript:popupReportDialog(4, 189275);" class="btn_report_error writeAction"> </a></li>
<li>
<a href="https://www.metal-archives.com/bands/Fauna/121144" title="Back to Fauna" class="btn_back"> </a>
</li>
</ul>
</div>

<div id="album_info">

<h1 class="album_name"><a href="https://www.metal-archives.com/albums/Fauna/The_Hunt/189275">The Hunt</a></h1>
<h2 class="band_name">
<a href="https://www.metal-archives.com/bands/Fauna/121144">Fauna</a>
</h2>

<div class="clear block_spacer_5"></div>
<div id="message"> </div>

<div class="clear block_spacer_20"></div>
<dl class="float_left">
<dt>Type:</dt>
<dd>Full-length</dd>
<dt>Release date:</dt>
<dd>2007</dd>
<dt>Catalog ID:</dt>
<dd>N/A</dd>
<dt>Version desc.:</dt>
<dd>Limited edition</dd>
</dl>
<dl class="float_right">
<dt>Label:</dt>
<dd>Independent</dd>
<dt>Format:</dt>
<dd>CD</dd>
<dt>Reviews:</dt>
<dd>
1 <a href="https://www.metal-archives.com/reviews/Fauna/The_Hunt/189275/">review</a> (avg. 85%)
</dd>
</dl>
</div>
<div id="album_tabs" class="clear tabs block_spacer_top_20 ui-tabs no-js">
<ul class="ui-tabs-nav">
<li><a href="#album_tabs_tracklist">Songs</a></li>
<li><a href="#album_tabs_lineup">Lineup</a></li>
<li><a href="https://www.metal-archives.com/release/ajax-versions/current/189275/parent/189275">Other versions</a></li> <li><a href="#album_tabs_reviews">Reviews</a></li> <li><a href="#album_tabs_notes">Additional notes</a></li> </ul>

<div id="album_tabs_tracklist" class="ui-tabs-hide">
<div id="album_songs" class="tabs2lvl">

</div>
<div class="ui-tabs-panel-content block_spacer_top_20">
<script type="fa3d21b0fdc73efe12102b0f-text/javascript">
function toggleLyrics(songId) {
    var lyricsRow = $('#song' + songId);
    lyricsRow.toggle();
    var lyrics = $('#lyrics_' + songId);
	if (lyrics.html() == '(loading lyrics...)') {
    	var realId = songId;
		if(!$.isNumeric(songId.substring(songId.length -1, songId.length))) {
			realId = songId.substring(0, songId.length -1);
		}
		lyrics.load(URL_SITE + "release/ajax-view-lyrics/id/" + realId);
    }
    // toggle link
    var linkLabel = "lyrics";
    $("#lyricsButton" + songId).text(lyricsRow.css("display") == "none" ? "Show " + linkLabel : "Hide " + linkLabel);
    return false;
}

</script>
<table class="display table_lyrics" cellpadding="0" cellspacing="0">
<tbody>
<tr class="even">
<td width="20"><a name="1340426" class="anchor"> </a>1.</td>
<td class="wrapWords">
The Hunt
</td>
<td align="right">01:19:57</td>
<td nowrap="nowrap">&nbsp;
<a id="lyricsButton1340426" href="#1340426" onclick="if (!window.__cfRLUnblockHandlers) return false; toggleLyrics('1340426'); return false;" data-cf-modified-fa3d21b0fdc73efe12102b0f-="">Show lyrics</a>
</td>
</tr>
<tr id="song1340426" class="displayNone" height="0">
<td>&nbsp;</td>
<td colspan="3" id="lyrics_1340426">(loading lyrics...)</td>
</tr>
<tr>
<td colspan="2">&nbsp;</td>
<td align="right"><strong>01:19:57</strong></td>
<td>&nbsp;</td>
</tr>
</tbody>
</table>
</div>
</div>


<div id="album_tabs_lineup">
<div id="album_members" class="tabs2lvl">
<ul>
<li><a href="#album_members_lineup">Band members</a></li> </ul>

<div id="album_members_lineup">
<div class="ui-tabs-panel-content">
<table class="display lineupTable" cellpadding="0" cellspacing="0">
<tr class="lineupRow">
<td width="300" valign="top">
<a href="https://www.metal-archives.com/artists/Vines/58139">Vines</a>
</td>
<td>
Drums, Vocals </td>
</tr>
<tr class="lineupRow">
<td width="300" valign="top">
<a href="https://www.metal-archives.com/artists/Echtra/58138">Echtra</a>
</td>
<td>
Guitars, Vocals </td>
</tr>
</table>
</div>
</div>
</div>
</div>





<div id="album_tabs_reviews" class="ui-tabs-hide">
<div id="album_reviews" class="tabs2lvl">
<div class="tool_strip top right writeAction">
<ul>
<li><a href="https://www.metal-archives.com/review/write/releaseId/189275" class="btn_add">Add</a></li>
</ul>
</div>
</div>
<div class="ui-tabs-panel-content block_spacer_top_36">
<table id="review_list" class="display" cellpadding="0" cellspacing="0">
<tr class="odd">
<td nowrap="nowrap"><a href="https://www.metal-archives.com/reviews/Fauna/The_Hunt/189275/drengskap/107392" title="Read" class="iconContainer ui-state-default ui-corner-all"><span class="ui-icon ui-icon-search">Read</span></a></td>
<td>Hunting and gathering</td>
<td>85%</td>
<td><a href="https://www.metal-archives.com/users/drengskap" class="profileMenu">drengskap</a></td>
<td>February 28th, 2010</td>
</tr>
</table>
</div>
</div>


<div id="album_tabs_notes" class="ui-tabs-hide">
<div id="album_notes" class="tabs2lvl"></div>
<div class="ui-tabs-panel-content  block_spacer_top_20">
<p class="block_spacer_20">CD in silkscreened black on black paper. This album was self-released, but should not be considered a demo.<br />
<br />
Re-released in December 2009 by Aurora Borealis and Stellar Auditorium, remastered by Mell Detmer (Earth, Sunn O))) etc) and with live drums replacing the original electronic drums.<br />
<br />
The re-released version is sequenced as seven tracks with the following titles and running times:<br />
1. The Door 05:43 (instrumental)<br />
2. Hunger 08:33<br />
3. Setting Out 11:23<br />
4. The Scent 17:32<br />
5. Nocturne 04:45<br />
6. Tracking 09:07<br />
7. The Kill... Fulfillment 14:44</p>
</div>
</div>

</div>

<div id="auditTrail">
<table>
<tr>
<td>Added by: <a href="https://www.metal-archives.com/users/pankkake" class="profileMenu">pankkake</a></td>
<td align="right">Modified by: <a href="https://www.metal-archives.com/users/Sodomatic_Slaughter" class="profileMenu">Sodomatic_Slaughter</a></td>
</tr>
<tr>
<td>Added on: 2008-03-08 09:13:13</td>
<td align="right">Last modified on: 2018-04-05 08:01:24</td>
</tr>
<tr>
<td valign="top">
&nbsp;
</td>
<td align="right" valign="top">
</td>
</tr>
</table>
</div>

</div>
</div>
</div>
<script src="https://ajax.cloudflare.com/cdn-cgi/scripts/7089c43e/cloudflare-static/rocket-loader.min.js" data-cf-settings="fa3d21b0fdc73efe12102b0f-|49" defer=""></script></body>
</html>
	`))}}}

	tracks, cover, err := GetAlbumInfo(client, albumData)

	if err != nil {
		t.Errorf("TestLongerTracks shouldn't fail.")
	}

	if len(tracks) != 1 {
		t.Errorf("The Hunt by Fauna has only 1 track not %d.", len(tracks))
	}

	if tracks[0].Name != "The Hunt" {
		t.Errorf("The Hunt by Fauna first track is called 'The Hunt', not '%s'.", tracks[0].Name)
	}

	if tracks[0].Hours != 1 {
		t.Errorf("The Hunt by Fauna hours should be 1, not '%d'.", tracks[0].Hours)
	}

	if tracks[0].Minutes != 19 {
		t.Errorf("The Hunt by Fauna minutes should be 5, not '%d'.", tracks[0].Minutes)
	}

	if tracks[0].Seconds != 57 {
		t.Errorf("The Hunt by Fauna seconds should be 57, not '%d'.", tracks[0].Seconds)
	}

	if cover != "https://www.metal-archives.com/images/1/8/9/2/189275.jpg" {
		t.Errorf("The Hunt by Fauna has cover located in 'https://www.metal-archives.com/images/1/8/9/2/189275.jpg', not %s'.", cover)
	}

}
