package ui

import (
	"encoding/hex"
	"image"
)

const (
	antTexPixEncoded = "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000021212102222222570000000000000000000000000000000000000000000000000000000000000000000000002222220022222258222222030000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002222226d2222228d000000000000000000000000000000000000000000000000000000000000000000000000000000002222228b2222226e1b1b1b000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222226222222f4222222420000000000000000000000000000000000000000000000000000000000000000000000000000000022222240222222f42222222600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222201222222b5222222e0222222050000000000000000000000000000000000000000000000000000000000000000000000000000000023232305222222e0222222b42121210100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002222224c222222fe2222227b1e1e1e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002222227a222222fe2222224a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222202222222ce222222e82222221200000000000000000000000000000000222222032222222622222223232323020000000000000000000000000000000022222212222222e9222222cc2323230300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222229222222ff2222228e2222220f2525250000000000282828002a2a2a3d363636d3282828fe222222fe222222d22222223c00000000000000001f1f1f002222220e2222228e222222fe222222280000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002222220d222222a9222222fe222222f12222229422222217262626434e4e4ef5515151ff383838ff222222ff222222ff222222f5222222422222221822222294222222f1222222fe222222ab2222220e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000202020012222224c222222de222222ff222222e93c3c3ceb5e5e5eff515151ff3b3b3bff222222ff222222ff232323ff222222e9222222e9222222ff222222df2222224f22222201000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002222220d22222284222222f55c5c5cff5e5e5eff515151ff3a3a3aff222222ff222222ff242424ff292929ff222222f7222222862222220e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222206212121000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002222220b353535eb6c6c6cff5e5e5eff515151ff393939ff222222ff222222ff242424ff383838ff222222ed2222220c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020202000232323061b1b1b000000000000000000000000002222227a2222225c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002222224c515151ff6c6c6cff5e5e5eff515151ff393939ff222222ff222222ff242424ff3f3f3fff2d2d2dff2222224e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000222222582222227d1919190000000000000000000000000022222214222222e12222225222222200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002222228e686868ff6c6c6cff5e5e5eff515151ff383838ff222222ff222222ff252525ff3f3f3fff414141ff2222229000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002222224f222222e122222214000000000000000000000000000000000000000022222251222222f22222223500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000282828be757575ff6c6c6cff5e5e5eff515151ff383838ff222222ff222222ff252525ff3f3f3fff525252ff222222c10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222233222222f222222254000000000000000000000000000000000000000000000000212121012222229a222222de22222214000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222200323232df787878ff6c6c6cff5e5e5eff515151ff393939ff222222ff222222ff252525ff3f3f3fff5d5d5dff262626e122222200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222212222222dc2222229a22222200000000000000000000000000000000000000000000000000000000002222220b222222d0222222b02222220f000000000000000000000000000000000000000000000000000000000000000000000000000000001f1f1f02393939ef787878ff6c6c6cff5e5e5eff515151ff383838ff222222ff222222ff252525ff3f3f3fff606060ff2d2d2df121212102000000000000000000000000000000000000000000000000000000000000000000000000000000002222220e222222ae222222d02222220b00000000000000000000000000000000000000000000000000000000000000000000000022222227222222ee222222e12222225023232301000000000000000000000000000000000000000000000000000000000000000022222202292929f06b6b6bff6c6c6cff5e5e5eff515151ff303030ff222222ff222222ff252525ff3f3f3fff606060ff313131f3212121030000000000000000000000000000000000000000000000000000000000000000232323012222224f222222df222222ef22222228000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222241222222f4222222fe222222ab2222221b000000000000000000000000000000000000000000000000000000002d2d2d00222222be242424ff303030ff323232ff2c2c2cff222222ff222222ff222222ff252525ff3e3e3eff494949ff242424c01e1e1e00000000000000000000000000000000000000000000000000000000002222221a222222ab222222fe222222f4222222420000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222240222222eb222222ff222222ec222222632323230200000000000000000000000000000000000000000000000022222220222222a3222222d8222222f6222222ff222222ff222222ff222222ff222222f6232323d8222222a3222222210000000000000000000000000000000000000000000000002222220322222265222222ec222222ff222222ec2222224000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222220222222b8222222fe222222ff222222c32222224f222222070000000000000000000000000000000000000000222222052f2f2f8a4d4d4dfa4d4d4dff292929ff222222ff222222ff222222fa2222228e2222220600000000000000000000000000000000000000002222220722222251222222c5222222ff222222fe222222ba222222210000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000212121022222224a222222bc222222fd222222ff222222e022222279222222130000000000000000232323012e2e2e97606060ff5e5e5eff515151ff353535ff222222ff222222ff242424ff252525ff2222229b2121210100000000000000002222221222222279222222e0222222ff222222fd222222be2222224d2222220200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002323230022222226222222c8222222ff222222ff222222f022222283222222112222223c555555fd6c6c6cff5e5e5eff515151ff383838ff222222ff222222ff252525ff3b3b3bff282828fe2222223e2222221022222281222222ef222222ff222222ff222222c8222222282222220000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222214222222be222222ff222222ff222222ff222222e7252525ca727272ff6c6c6cff5e5e5eff515151ff383838ff222222ff222222ff252525ff3f3f3fff464646ff222222ca222222e6222222ff222222ff222222ff222222be22222215000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000212121042222226f222222e7222222ff222222ff2e2e2eff787878ff6c6c6cff5e5e5eff515151ff373737ff222222ff222222ff252525ff3f3f3fff5b5b5bff242424ff222222ff222222ff222222e722222271222222050000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002222220b222222552222229c343434f3787878ff6c6c6cff5e5e5eff515151ff363636ff222222ff222222ff252525ff3f3f3fff606060ff2b2b2bf42222229d222222562222220b212121000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002a2a2a00222222002a2a2a000000000000000000000000001f1f1f00363636cc787878ff6c6c6cff5e5e5eff515151ff363636ff222222ff222222ff252525ff3f3f3fff606060ff313131ce2b2b2b000000000000000000000000000000000025252500232323000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000222222002222220a2222223d22222275222222a4222222b6222222a32222227d2222224b2222221720202000323232be787878ff6c6c6cff5e5e5eff515151ff373737ff222222ff222222ff252525ff3f3f3fff606060ff2f2f2fbe20202000222222172222224b2222227c222222a3222222b7222222a5222222772222223e2222220b2b2b2b0000000000000000000000000000000000000000000000000000000000000000000000000028282800242424002222220022222200212121022222221d2222226c222222b7222222f2222222ff222222ff222222ff222222ff222222ff222222ff222222ff222222fb222222d2292929e3777777ff6c6c6cff5e5e5eff515151ff373737ff222222ff222222ff252525ff3f3f3fff606060ff292929e3222222d3222222fb222222ff222222ff222222ff222222ff222222ff222222ff222222ff222222f2222222b72222226c2222221c22222201222222001c1c1c002121210019191900000000002222225f222222b2222222c7222222d1222222d8222222e1222222fa222222ff222222ff222222fe222222f6222222e0222222c2222222a5222222a5222222e0222222fe222222ff222222ff232323ff6f6f6fff6c6c6cff5e5e5eff515151ff373737ff222222ff222222ff252525ff3f3f3fff5c5c5cff242424ff222222ff222222ff222222fe222222e2222222a9222222a8222222c6222222e3222222f8222222ff222222ff222222ff222222fa222222e1222222d8222222d0222222c5222222b02222225f22222205222222222222224022222251222222592222224f2222224222222242222222342222221e2222220a23232302212121000000000000000000222222032222222d22222275222222be222222f7606060ff6c6c6cff5e5e5eff515151ff373737ff222222ff222222ff252525ff3f3f3fff535353ff222222f6222222bc222222742222222d23232304333333000000000021212100222222022222220c22222221222222372222224522222244222222522222225c2222225422222244222222252222220600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000027272700222222394a4a4aff6c6c6cff5e5e5eff515151ff363636ff222222ff222222ff252525ff3f3f3fff464646ff22222237000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222205323232df6b6b6bff5f5f5fff515151ff353535ff222222ff222222ff252525ff3f3f3fff383838df22222205000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000222222764e4e4eff5f5f5fff515151ff323232ff222222ff222222ff252525ff3e3e3eff29292977000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000021212100222222112222224b232323f73b3b3bff434343ff282828ff222222ff222222ff252525ff2c2c2cf7222222492222220f24242400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002121210222222247222222b3222222f8222222ff222222ff222222ff222222ff222222ff222222ff222222ff222222ff222222ff222222ff222222f7222222b122222244222222020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002121210022222231222222c0222222fe222222ff222222ff222222ff222222ff222222ff222222ff222222ff222222ff222222ff222222ff222222ff222222ff222222ff222222ff222222fe222222be2222222f222222000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002222220522222281222222f9222222ff222222ff222222ff222222f0222222fe262626ff303030ff2f2f2fff232323ff222222ff222222ff222222ff222222ff222222ff222222ef222222ff222222ff222222ff222222f92222227e212121050000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222211222222b8222222ff222222ff222222ff222222ca22222273242424e7525252ff696969ff5f5f5fff515151ff333333ff222222ff222222ff232323ff272727ff222222ff222222e522222272222222c9222222ff222222ff222222ff222222b62222220e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222213222222cc222222ff222222fe222222b82222224022222202222222ac4f4f4fff787878ff6c6c6cff5f5f5fff515151ff373737ff222222ff222222ff252525ff404040ff4c4c4cff262626ff222222aa212121012222223f222222b7222222fe222222ff222222c72222221200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222213222222cb222222ff222222f822222262222222010000000022222216242424f9737373ff787878ff6c6c6cff5f5f5fff515151ff373737ff222222ff222222ff252525ff404040ff616161ff525252ff222222f822222215000000002222220122222260222222f7222222ff222222c8222222120000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222213222222cc222222ff222222e22222223d2222220000000000000000002222223d2f2f2fff7a7a7aff787878ff6c6c6cff5f5f5fff515151ff373737ff222222ff222222ff252525ff404040ff616161ff727272ff232323ff2222223c0000000000000000232323002222223c222222e1222222ff222222c722222212000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222212222222ca222222ff222222b4222222170000000000000000000000000000000022222243363636ff7a7a7aff787878ff6c6c6cff5f5f5fff515151ff373737ff222222ff222222ff252525ff404040ff616161ff797979ff2d2d2dff222222440000000000000000000000000000000022222216222222b5222222ff222222c82222221200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222210222222c9222222fa2222227821212103000000000000000000000000000000000000000022222237383838ff7a7a7aff787878ff6c6c6cff5f5f5fff515151ff373737ff222222ff222222ff252525ff404040ff616161ff797979ff343434ff2222223800000000000000000000000000000000000000002222220322222279222222f9222222c52222220f000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002222220b222222c3222222e32222223f2020200000000000000000000000000000000000000000000000000022222220353535fe7a7a7aff787878ff6c6c6cff5f5f5fff515151ff373737ff222222ff222222ff252525ff404040ff616161ff797979ff353535ff222222200000000000000000000000000000000000000000000000000000000022222241222222e3222222bf2222220a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000202020002222228b222222c92222221a0000000000000000000000000000000000000000000000000000000000000000232323082a2a2af0797979ff787878ff6c6c6cff5f5f5fff515151ff373737ff222222ff222222ff252525ff404040ff616161ff797979ff303030f02222220800000000000000000000000000000000000000000000000000000000000000002222221b222222cb2222228a1f1f1f000000000000000000000000000000000000000000000000000000000000000000000000000000000022222206222222e42222222100000000000000000000000000000000000000000000000000000000000000000000000024242400222222c06c6c6cff787878ff6c6c6cff5f5f5fff515151ff373737ff222222ff222222ff252525ff404040ff616161ff767676ff262626c12121210000000000000000000000000000000000000000000000000000000000000000000000000022222224222222e522222206000000000000000000000000000000000000000000000000000000000000000000000000000000002222222e222222c82222220000000000000000000000000000000000000000000000000000000000000000000000000000000000222222754e4e4eff787878ff6c6c6cff5f5f5fff515151ff363636ff222222ff222222ff252525ff404040ff616161ff626262ff222222770000000000000000000000000000000000000000000000000000000000000000000000000000000021212100222222ca2222222d0000000000000000000000000000000000000000000000000000000000000000000000000000000022222251222222a200000000000000000000000000000000000000000000000000000000000000000000000000000000000000002222221a2a2a2af6727272ff6c6c6cff5f5f5fff515151ff333333ff222222ff222222ff252525ff404040ff616161ff424242f72222221b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000222222a32222225100000000000000000000000000000000000000000000000000000000000000000000000000000000222222642222229200000000000000000000000000000000000000000000000000000000000000000000000000000000000000001d1d1d002222228f404040ff6c6c6cff5f5f5fff515151ff2a2a2aff222222ff222222ff252525ff404040ff555555ff2626269200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222294222222650000000000000000000000000000000000000000000000000000000000000000000000000000000022222269222222920000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002222220f222222d63a3a3aff525252ff3c3c3cff222222ff222222ff222222ff252525ff3e3e3eff2e2e2ed922222210000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000222222932222226a00000000000000000000000000000000000000000000000000000000000000000000000000000000222222642222229e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222224222222dc222222ff222222ff222222ff222222ff222222ff232323ff262626df22222226000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002222229f222222650000000000000000000000000000000000000000000000000000000000000000000000000000000022222253222222b000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002222221322222297222222f4222222ff222222ff222222f52222229c222222150000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000222222b0222222520000000000000000000000000000000000000000000000000000000000000000000000000000000022222234222222c424242400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000022222200222222102222223c2222223f2222221200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000021212100222222c422222235000000000000000000000000000000000000000000000000000000000000000000000000000000002222220c222222cd232323000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000021212100222222ce2222220c00000000000000000000000000000000000000000000000000000000000000000000000000000000222222002222227223232300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002020200022222274212121000000000000000000000000000000000000000000"
)

var (
	antTex *image.NRGBA
)

func init() {
	// r, _ := os.Open("./ui/ant.png")
	// img, _ := png.Decode(r)
	// rgba := img.(*image.NRGBA)
	// println(hex.EncodeToString(rgba.Pix))
	// println(rgba.Stride)
	// println(rgba.Rect)

	pix, _ := hex.DecodeString(antTexPixEncoded)
	antTex = &image.NRGBA{
		Pix:    []uint8(pix),
		Stride: 200,
		Rect:   image.Rect(0, 0, 50, 55),
	}
}
