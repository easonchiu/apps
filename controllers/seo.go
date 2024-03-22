package controllers

import "github.com/gin-gonic/gin"

func (Controller) PageRobots(ctx *gin.Context) {
	ctx.String(200, "Sitemap: https://www.ysgamestudio.com/sitemap.xml")
}

func (Controller) PageSitemap(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/xml")
	// 构建Sitemap数据
	sitemap := `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
        <loc>https://www.ysgamestudio.com</loc>
    </url>
    <url>
        <loc>https://www.ysgamestudio.com/games/sudoku-crown</loc>
    </url>
    <url>
        <loc>https://www.ysgamestudio.com/games/digit-merge</loc>
    </url>
    <url>
        <loc>https://www.ysgamestudio.com/games/block-cuties</loc>
    </url>
    <url>
        <loc>https://www.ysgamestudio.com/privacy</loc>
    </url>
    <url>
        <loc>https://www.ysgamestudio.com/contact</loc>
    </url>
</urlset>`

	// 返回XML响应
	ctx.String(200, sitemap)
}

func (Controller) PageAppADs(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/plain")
	ctx.Header("Cache-Control", "max-age=3600")
	txt := `152media.info, 152M324, RESELLER
adcolony.com, 0a0f72cd66122f31, RESELLER, 1ad675c9de6b5176
adcolony.com, 801e49d1be83b5f9, RESELLER, 1ad675c9de6b5176
adelement.com, 30167, RESELLER
adform.com, 2742, RESELLER
admanmedia.com, 613, RESELLER
admixer.net, 4ca083cd-412f-49fe-b832-5b66ee490d9a, RESELLER
admixplay.com, 2012, RESELLER
adsyield.com, 1233, RESELLER
adview.com, 48045325, RESELLER, 1b2cc038a11ea319
algorix.co, 54616, RESELLER, 5b394c12fea27a1d
algorix.co, 60016, RESELLER, 5b394c12fea27a1d
appads.in, 106246, RESELLER
appads.in, 107606, RESELLER
appnexus.com, 10617, RESELLER, f5ab79cb980f11d1
appnexus.com, 11826, RESELLER
appnexus.com, 11924, RESELLER, f5ab79cb980f11d1
appnexus.com, 12976, RESELLER, f5ab79cb980f11d1
appnexus.com, 13227, RESELLER
appnexus.com, 13293, RESELLER, f5ab79cb980f11d1
appnexus.com, 13297, RESELLER, f5ab79cb980f11d1
appnexus.com, 2764, RESELLER, f5ab79cb980f11d1
appnexus.com, 7353, RESELLER, f5ab79cb980f11d1
appnexus.com, 9393, RESELLER, f5ab79cb980f11d1
bidease.com, bidease_seller_20, RESELLER
bidmachine.io, 113, RESELLER
bidmachine.io, 114, RESELLER
bidmachine.io, 162, RESELLER
bidmachine.io, 67, RESELLER
contextweb.com, 562499, RESELLER, 89ff185a4c4e857c
contextweb.com, 562791, RESELLER, 89ff185a4c4e857c
contextweb.com, 562842, RESELLER, 89ff185a4c4e857c
contextweb.com, 562944, RESELLER, 89ff185a4c4e857c
criteo.com, B-057601, RESELLER, 9fac4a4a87c2a44f
criteo.com, B-057955, RESELLER, 9fac4a4a87c2a44f
criteo.com, B-063105, RESELLER, 9fac4a4a87c2a44f
criteo.com, B-072395, RESELLER, 9fac4a4a87c2a44f
exchange.admazing.co, AM3, RESELLER
google.com, pub-3481503181298547, RESELLER
google.com, pub-3481503181298547, RESELLER, f08c47fec0942fa0
improvedigital, 2343, RESELLER
improvedigital.com, 1604, RESELLER
improvedigital.com, 2052, RESELLER
indexexchange.com, 186081, RESELLER, 50b1c356f2c5c8fc
indexexchange.com, 186684, RESELLER, 50b1c356f2c5c8fc
indexexchange.com, 194080, RESELLER, 50b1c356f2c5c8fc
indexexchange.com, 194730, RESELLER, 50b1c356f2c5c8fc
indexexchange.com, 198417, RESELLER, 50b1c356f2c5c8fc
indexexchange.com, 199195, RESELLER, 50b1c356f2c5c8fc
inmobi.com, 062029933580429f9920bad476d8d70a, RESELLER, 83e75a7ae333ca9d
inmobi.com, 3a4f7da341dd490cbb7dde02b126275e, RESELLER, 83e75a7ae333ca9d
inmobi.com, 61d733c3779d43e590c51c8bc078e10c, RESELLER, 83e75a7ae333ca9d
inmobi.com, 7847fe1f9ac54b4abe609cde4011243b, RESELLER, 83e75a7ae333ca9d
kidoz.net, 14260, RESELLER, a109366414b7335e
lemmatechnologies.com, 362, RESELLER, 7829010c5bebd1fb
liftoff.io, 7f6945815e6, RESELLER
loopme.com, 11414, RESELLER, 6c8d5f95897a5a3b
loopme.com, 11635, RESELLER, 6c8d5f95897a5a3b
loopme.com, 9621, RESELLER, 6c8d5f95897a5a3b
lunamedia.io, 2fdb16a1-5025-4c41-9b17-e01e47ebc8b4, RESELLER, 524ecb396915caaf
meitu.com, 282, RESELLER
meitu.com, 699, RESELLER
mintegral.com, 10038, RESELLER
onetag.com, 7d9af0b85b5070e, RESELLER
openx.com, 540679900, RESELLER, 6a698e2ec38604c6
openx.com, 540838151, RESELLER, 6a698e2ec38604c6
openx.com, 540866936, RESELLER, 6a698e2ec38604c6
openx.com, 540871654, RESELLER, 6a698e2ec38604c6
openx.com, 541031350, RESELLER, 6a698e2ec38604c6
opera.com, pub6794361378752, RESELLER, 55a0c5fd61378de3
playwire.com, 1018531, RESELLER
playwire.com, 1025000, RESELLER
pokkt.com, 6246, RESELLER, c45702d9311e25fd
pokkt.com, 7606, RESELLER, c45702d9311e25fd
pubmatic.com, 156494, RESELLER, 5d62403b186f2ace
pubmatic.com, 156520, RESELLER, 5d62403b186f2ace
pubmatic.com, 157384, RESELLER, 5d62403b186f2ace
pubmatic.com, 157559, RESELLER, 5d62403b186f2ace
pubmatic.com, 158060, RESELLER, 5d62403b186f2ace
pubmatic.com, 158481, RESELLER, 5d62403b186f2ace
pubmatic.com, 158565, RESELLER, 5d62403b186f2ace
pubmatic.com, 159501, RESELLER, 5d62403b186f2ace
pubmatic.com, 159906, RESELLER, 5d62403b186f2ace
pubmatic.com, 160113, RESELLER, 5d62403b186f2ace
pubmatic.com, 160295, RESELLER, 5d62403b186f2ace
pubmatic.com, 160456, RESELLER, 5d62403b186f2ace
pubmatic.com, 160493, RESELLER, 5d62403b186f2ace
pubmatic.com, 160536, RESELLER, 5d62403b186f2ace
pubmatic.com, 160846, RESELLER, 5d62403b186f2ace
pubmatic.com, 160974, RESELLER, 5d62403b186f2ace
pubmatic.com, 161018, RESELLER, 5d62403b186f2ace
pubmatic.com, 161136, RESELLER, 5d62403b186f2ace
pubmatic.com, 161344, RESELLER, 5d62403b186f2ace
pubmatic.com, 161372, RESELLER, 5d62403b186f2ace
pubmatic.com, 161853, RESELLER, 5d62403b186f2ace
pubmatic.com, 162223, RESELLER, 5d62403b186f2ace
pubmatic.com, 162239, RESELLER, 5d62403b186f2ace
pubmatic.com, 162882, RESELLER, 5d62403b186f2ace
pubmatic.com, 162968, RESELLER, 5d62403b186f2ace
pubmatic.com, 163319, RESELLER, 5d62403b186f2ace
pubmatic.com, 163420, RESELLER, 5d62403b186f2ace
pubnative.net, 1003601, RESELLER, d641df8625486a7b
pubnative.net, 1006955, RESELLER, d641df8625486a7b
pubnative.net, 1007040, RESELLER, d641df8625486a7b
pubnative.net, 1007063, RESELLER, d641df8625486a7b
pubnative.net, 1007262, RESELLER, d641df8625486a7b
pubnative.net, 1007501, RESELLER, d641df8625486a7b
pubnative.net, 1008770, RESELLER, d641df8625486a7b
reforge.in, 107, RESELLER
reforge.in, 353, RESELLER
rhythmone.com, 3383599585, RESELLER, a670c89d4a324e47
rhythmone.com, 3630748062, RESELLER, a670c89d4a324e47
risecodes.com, 6486c6155b231000010244b0, RESELLER
rubiconproject.com, 12556, RESELLER, 0bfd66d529a55807
rubiconproject.com, 14334, RESELLER, 0bfd66d529a55807
rubiconproject.com, 15854, RESELLER, 0bfd66d529a55807
rubiconproject.com, 17608, RESELLER, 0bfd66d529a55807
rubiconproject.com, 18202, RESELLER, 0bfd66d529a55807
rubiconproject.com, 20014, RESELLER, 0bfd66d529a55807
rubiconproject.com, 20050, RESELLER, 0bfd66d529a55807
rubiconproject.com, 20086, RESELLER, 0bfd66d529a55807
rubiconproject.com, 20744, RESELLER, 0bfd66d529a55807
rubiconproject.com, 21526, RESELLER, 0bfd66d529a55807
rubiconproject.com, 23876, RESELLER, 0bfd66d529a55807
rubiconproject.com, 24170, RESELLER, 0bfd66d529a55807
rubiconproject.com, 24362, RESELLER, 0bfd66d529a55807
rubiconproject.com, 24400, RESELLER, 0bfd66d529a55807
rubiconproject.com, 24526, RESELLER, 0bfd66d529a55807
rubiconproject.com, 24600, RESELLER, 0bfd66d529a55807
rubiconproject.com, 26184, RESELLER, 0bfd66d529a55807
rubiconproject.com, 26386, RESELLER, 0bfd66d529a55807
sharethrough.com, bXzzb1bR, RESELLER, d53b998a7bd4ecd2
sharethrough.com, iBAzay96, RESELLER, d53b998a7bd4ecd2
smaato.com, 1100004890, RESELLER, 07bcf65f187117b4
smaato.com, 1100055750, RESELLER, 07bcf65f187117b4
smaato.com, 1100056144, RESELLER, 07bcf65f187117b4
smartclip.net, 13564, RESELLER
spotx.tv, 117872, RESELLER, 7842df1d2fe2db34
spotx.tv, 149886, RESELLER, 7842df1d2fe2db34
spotx.tv, 71426, RESELLER, 7842df1d2fe2db34
spotxchange.com, 117872, RESELLER, 7842df1d2fe2db34
spotxchange.com, 149886, RESELLER, 7842df1d2fe2db34
spotxchange.com, 71426, RESELLER, 7842df1d2fe2db34
thebrave.io, 1234594, RESELLER, c25b2154543746ac
thebrave.io, 1234634, RESELLER, c25b2154543746ac
thebrave.io, 1234678, RESELLER, c25b2154543746ac
thebrave.io, 9840732, RESELLER, c25b2154543746ac
themediagrid.com, FH3TKJ, RESELLER, 35d5010d7789b49d
themediagrid.com, R28I9J, RESELLER, 35d5010d7789b49d
themediagrid.com, SWH94X, RESELLER, 35d5010d7789b49d
toponad.com, 165535463cfb70, RESELLER, 1d49fe424a1a456d
tredio.io, 357a6fdf7642bf815a88822c447d9dc433546, RESELLER
triplelift.com, 13883, RESELLER, 6c33edb13117fd86
unity.com, 4746947, DIRECT, 96cabb5fbdde37a7
velismedia.com, 1137, RESELLER
velismedia.com, 725, RESELLER
video.unrulymedia.com, 123476257, RESELLER
video.unrulymedia.com, 170071695, RESELLER
video.unrulymedia.com, 2464975885, RESELLER
video.unrulymedia.com, 3383599585, RESELLER, a670c89d4a324e47
video.unrulymedia.com, 3630748062, RESELLER, a670c89d4a324e47
video.unrulymedia.com, 392157850, RESELLER
video.unrulymedia.com, 819070080, RESELLER
video.unrulymedia.com, 995298565, RESELLER
videoheroes.tv, 212499, RESELLER, 064bc410192443d8
vidoomy.com, 7646534, RESELLER
webeyemob.com, 70097, RESELLER
webeyemob.com, 70104, RESELLER
xandr.com, 13167, RESELLER, f5ab79cb980f11d1
xandr.com, 13297, RESELLER, f5ab79cb980f11d1
xandr.com, 13799, RESELLER
xandr.com, 14082, RESELLER
xandr.com, 15769, RESELLER
xandr.com, 7353, RESELLER, f5ab79cb980f11d1
xandr.com, 8804, RESELLER, f5ab79cb980f11d1
xapads.com, 117219, RESELLER
xapads.com, 155860, RESELLER
yandex.com, 40427798, RESELLER
yandex.com, 95746987, RESELLER
yandex.com, 97539269, RESELLER
yeahmobi.com, 5135195, RESELLER`

	// 返回
	ctx.String(200, txt)
}
