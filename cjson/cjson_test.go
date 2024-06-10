package cjson_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/0x587/go-cjson/cjson"
)

var table = []string{
	`{"string":"Hello, World!","number":42,"float":3.14159,"boolean":true,"null":null,"object":{"nestedString":"Nested Hello","nestedNumber":100,"nestedArray":[1,2,3],"nestedObject":{"deeplyNestedString":"Deep"}},"array":["String in array",123,false,null,{"key":"value"},[1,2,3]]}`,
	`{"crop":"east","neighbor":-1360140019,"add":-2090583717,"proud":"alike","leader":false,"construction":true,"friendly":454682939.3501859,"function":true}`,
	`{"behavior":{"frequently":false,"ever":{"closer":{"pack":"barn","deeply":{"beauty":["stove",-1895154696.8433795,"repeat","truck","copper",{"rapidly":true,"active":[463789631.3256922,"grandmother",false,-1452973551,"across",1673062295.0531921,-2038167605.1869602,"describe",-1642543913.709558,true],"combine":"loss","during":true,"however":true,"not":22638349.960696697,"location":631647675,"it":"theory","will":"coal","slabs":true},"prize","use","dinner",true],"dinner":2035296494.303656,"use":"whatever","ourselves":false,"bear":"limited","map":false,"bank":1861462827,"log":1828093204,"oldest":"fallen","shaking":-1917002130},"daughter":"wheat","sail":"feathers","itself":"machinery","pink":"scientific","movement":true,"question":true,"driven":"national","bank":false},"breath":true,"beautiful":true,"strength":true,"report":1518610908.8946505,"provide":true,"information":"if","him":false,"social":126971080,"start":true},"whispered":true,"owner":"threw","foreign":true,"national":true,"did":false,"weather":"boy","west":true,"minute":true},"shut":"meal","winter":true,"state":"drive","sick":true,"baseball":false,"sand":-1289406717,"fire":"smaller","built":2017083043.5495548,"telephone":-1619931714.290587}`,
	`[[[{"heavy":"far","title":"him","whistle":"jack","climate":{"around":"frighten","back":false,"hope":[{"doing":{"phrase":"label","sold":["cowboy","map",false,{"gain":true,"lamp":"tune","pen":1375759840.8272104,"substance":"well","official":-313980616.92627144,"load":"rope","eaten":"protection","clay":1324392805,"shade":"became","ball":234984040.6118827},false,false,"pressure","likely",true,-2045187199.1055017],"life":-1779006125,"actual":-1498518238,"meet":-1657460791.4719086,"it":"unhappy","to":"slow","needle":false,"weight":"start","nice":"needle"},"well":true,"equal":959258587,"clock":"rush","boy":true,"please":"hot","skill":"air","dozen":false,"prevent":-1268630314.886057,"coal":364952255},2097553987,-422779125.5704336,true,"behavior","proud",-1337995023.441771,-2033168355,"recent",false],"interest":true,"consonant":false,"describe":"enough","double":1081620860,"wherever":false,"skill":"vegetable","jungle":"physical"},"chapter":true,"salt":true,"courage":true,"than":498911460,"speed":-1794141925,"proud":"soldier"},false,"leg",-1137035644.9649043,"recognize","learn",635375033,"discussion","shorter","came"],"pack","tank","glass","price",-651961283.8958282,"process",true,"more",2066923046],true,-274804747,20575303.257750988,"president",false,-255923660.66546345,false,"clearly",2075890685.513548]`,
	`["weigh",[false,"mixture","blood",{"stick":1588064606,"stood":true,"happened":{"principal":{"stranger":[{"busy":false,"fighting":[{"popular":[[{"indeed":"scientific","show":{"handle":["extra",-1748182755,false,248089149.9231305,false,true,-693008871.8127117,{"smoke":false,"ten":["question","talk",false,[[false,true,false,{"result":{"potatoes":"operation","team":"vertical","spread":-1176315860,"hour":false,"sky":"young","sweet":"wooden","finger":true,"like":-671785347.9864869,"those":-1055418367.3553848,"limited":true,"till":-409061123,"related":-804912861.3759727,"seldom":false,"handle":757886895,"cut":"grow","captured":"threw","cave":"door","blow":1543703783.1659708,"guess":1804168554},"pan":-330209540,"he":true,"generally":false,"concerned":"century","mistake":933630378,"wind":"fully","duck":-35431584.243386745,"combination":756156028.5871906,"slight":2024850426.0759516,"disease":false,"positive":true,"jar":-761762227.346261,"nor":"rhyme","spend":true,"electricity":"coach","poetry":"anyway","introduced":true,"coal":false,"enough":1266371217.6429062},false,true,1886864022.6729078,false,-1256689364,1897845169.1236348,true,-370987768.3448696,false,1723092860,"family",-558514667.7939725,false,false,"coffee",1816198376.4295096],"purpose",-443890147.7711129,2142212729,835992688.9388313,true,"remember",true,"second","limited",-798149667,"stone",-982277604.1934142,"kind","fruit",true,"might","usual",false,false],"none",-1179705020,true,"youth",true,-478093606,-191622426,true,"court",true,-85907021,"rich",false,"living",-148123433,false],"arm":false,"wave":-1662176830.171783,"pair":-986274141,"diagram":false,"combination":"ask","stood":-1962992676.144555,"explain":true,"function":true,"trace":"vertical","swept":434304716,"recognize":true,"essential":"partly","wood":false,"whose":true,"ship":"principle","locate":true,"sit":-1403652882.0512123,"watch":"parts"},"chest",1439972931,false,"ourselves","available","shirt",false,true,980868938.0110822,"desert",false,"health"],"naturally":false,"any":"industrial","luck":"stand","or":-300700869.23917675,"shoot":-1449060843,"thrown":"substance","evidence":"hello","thirty":true,"him":true,"might":-1034921902.815692,"level":2013185999.421616,"spirit":false,"apart":"division","upon":482732627.69560194,"blank":false,"slowly":true,"been":"circle","cost":"customs","other":-1409630961},"toy":"popular","alive":true,"spirit":true,"nature":true,"research":"degree","home":2113842600.392251,"monkey":-2133116289,"bone":"except","carefully":"mill","course":-858451179.1964054,"throughout":false,"particles":"children","whose":false,"fell":1733100140,"short":724757646,"motor":false,"park":222582433,"him":false},"clear",true,"loose","dollar",false,-432868419,false,-639025869.2674403,false,1667390499,false,-670771074.9752226,true,"figure",false,"refused",546904792,-1153462422,true],true,"discuss",1503411373.2068405,"refer","oxygen",false,"snow",false,"own","atom",true,"struck",270399695,607662651.9947343,true,-476258732.21068525,-1456812235,"family",-1745351130],"seldom":574129644,"together":false,"nine":1411575531.3381133,"cut":748877019,"met":-1480598037.0822563,"further":3460192.6280608177,"orange":"symbol","perhaps":"class","mill":"airplane","chamber":"battle","yourself":true,"shoot":1062489335,"cotton":1227154959,"smell":2027574094,"beauty":false,"short":1148265897.1233273,"frequently":"syllable","breakfast":"race","seeing":-2026598548},false,668825128.9304147,true,"station",true,-2080963011,false,"prove",true,true,true,-978136208.3206434,1751705292.3196182,2085339069.107658,false,1092338027,1854169361,1838506098,"after"],"up":true,"connected":-935440847.9878898,"join":"paragraph","themselves":false,"aloud":true,"molecular":false,"recognize":-404125993,"minute":471120272,"drink":1231397947,"climb":"chair","shown":-1325822704,"enough":1278276909,"ball":false,"solve":1711552770,"small":false,"joy":false,"curious":1849849373,"fill":579353694.4249377},-1805313693,"dark",true,false,true,127396738,false,"ten",1358782664.0505524,true,true,1146349737.080902,"direction",-670941424.0813642,-189939868.7801304,1642550321,false,-2082260838.3490705,313312146],"load":1971397514.8901052,"mail":"bite","written":-2022618874,"birds":true,"wagon":"valley","thread":"printed","goes":true,"pool":938668742,"lungs":1146941754,"daily":false,"coming":"eye","yellow":true,"metal":false,"frequently":true,"low":false,"meal":"us","good":496890363,"plural":true,"cause":"has"},"land":"recognize","gave":1002173880,"correctly":1786950050.0277567,"neighbor":-479900245,"terrible":true,"married":false,"contrast":"glad","shout":false,"vapor":"everyone","claws":"tiny","father":"entire","property":"softly","she":"brick","tube":false,"classroom":true,"source":"inch","dull":-752620583,"research":-488649777,"than":-473305042.862175},"science":-1248698823,"water":"specific","record":-1719497755.5249958,"weight":-820263308.1386251,"people":521102324,"dish":-1313136086,"process":1335511607,"island":true,"return":-66075893.72462416,"review":false,"state":-1076968343,"large":false,"anywhere":106435336.34556675,"breeze":-881476808.0496683,"steel":false,"sugar":"save","beautiful":35278682.30547428},false,"captain",538684531.4136987,true,"price","ring",1248934973,true,-1906136015,true,true,-1507479329.1192179,true,"empty",-873411900,947628285],-1714672994.2311335,"stepped",false,-1129844077.0743198,"few",true,false,587053526,779655222.9519467,"rabbit","weak","sudden",true,false,false,false,true,"produce"]`,
	`[[false,"slabs",["police",1143454774,"trip",2087133324.782494,{"sport":2028439637,"example":false,"atom":"anyone","becoming":[false,"adult",-902724308.3208666,1179025187,[{"arrow":"strong","thank":[1779287013.632563,true,{"wore":{"calm":1260371887.4991746,"form":["pack",false,[{"complete":{"funny":-322449419,"there":"development","stick":-307906845,"school":"stand","applied":["increase","gave",false,617330726.9734654,true,[1623028850,true,168363540.65181017,[[-1309104744.26132,false,false,{"task":{"pile":-952544674.3566437,"bring":{"add":"tired","shut":{"string":false,"who":"market","natural":"decide","cool":"teacher","sort":true,"forgotten":261346645.2675352,"load":1793531713,"liquid":"so","copper":true,"trunk":[{"balloon":-1979770450.9540477,"second":true,"cutting":[{"movie":{"shop":[{"five":true,"atmosphere":[true,true,-109573084.78938341,"nearby",true,"aloud",false,false,1524308120.5829787,"rapidly",true,-1577237192,false,"trap",true,"safety",1992402254,true,true,1901295482,false,false,-1982913950,"aid",796601444.080761,false,false,"life",853345034.4469428,"perfectly"],"start":false,"jack":-1448583503.8195186,"park":"something","shine":"plates","canal":false,"east":"repeat","telephone":false,"car":-475046426,"sides":"bit","swept":"which","taste":true,"on":false,"tell":1700808478,"accident":"labor","sell":-2078887592,"slowly":"anyway","variety":"every","matter":false,"get":false,"clear":809158581,"within":2077352017,"signal":true,"seed":false,"pour":-1940714620.212895,"climate":-960256334,"sweet":"trunk","felt":"explanation","branch":true},-1784169529.4815779,false,240040233.8589878,-645950360.489644,"food",1154522525.0722046,false,886393404.387362,"thick",308593899,135630708.82569456,true,"coal",false,1702985482,"written",true,false,1886789909,-1183845488.1779675,-1727548300,-991202412,-804929755.6075072,915745276,56118086.88983011,false,true,false,"test"],"unusual":true,"flow":-500485318.1237507,"mountain":false,"might":-514359195.2270684,"great":true,"correct":"forget","height":-1568926289.7601867,"railroad":"palace","gradually":"until","wherever":"swim","sister":false,"aside":"such","show":"outline","knew":true,"settlers":true,"forgotten":776063330.4155741,"guide":-429936861.5009284,"unless":true,"angle":2104799764,"letter":1505511669.8711495,"feathers":-1536865438,"high":"satellites","crowd":"process","glass":true,"highway":false,"floor":"sunlight","battle":"failed","that":true},"represent":1134032507,"cook":false,"accurate":991723319.7832808,"simplest":-483475332.0335264,"younger":"most","put":false,"against":true,"went":true,"mass":"say","camp":false,"classroom":2054453197.1552448,"colony":true,"per":true,"forty":false,"fine":true,"outer":-1004741047.0791488,"sentence":"route","being":false,"business":"stand","hungry":"possible","map":"street","various":true,"pack":true,"difficulty":"upward","mill":"best","city":-394851673,"spent":"second","stairs":"able"},"needed","moving",444926364,true,"claws",false,-1444044742,-2137993975,true,1348435996,"practical",-774177987.3548212,false,"cat",false,1236859892.978516,"ordinary",1749434409,"whether","ship","discussion",-774344345.5185542,1681585232,false,true,1517508666,"natural","train",740022141.3227544],"rose":"scared","fall":-345562224,"did":-619408394,"women":"letter","enough":"bus","cent":1614490684,"actual":852729257,"both":false,"kill":"swimming","atomic":true,"store":-470683917,"serious":false,"dozen":"pig","row":1383034853.8668585,"nobody":false,"when":"catch","people":false,"drew":"doing","chance":462867736.89075136,"any":-1503237223.8643022,"tales":-783642336.2241683,"industrial":true,"bicycle":"swimming","me":true,"ear":true,"understanding":-344791386,"pond":-137708456.77078533},false,true,1924010967,15084349.880126953,2078109108.8887267,false,"wish",499591582,"mass",false,1389145575.6530528,1480041804,"knife",true,"previous","pale","mission",-1355656785.9682288,true,"action",true,-1844042021.6345754,"attempt",628088938.6251187,"anywhere",false,false,-474333446,false],"serve":-1952898853.4472327,"stared":1114167222.095797,"bone":true,"camp":1127346173.2729092,"graph":false,"live":"run","do":"machine","two":true,"sport":true,"weak":true,"border":false,"hay":false,"molecular":1870073808.40905,"season":false,"useful":2007346551,"later":false,"plural":1688233509,"instance":217623832.64765835,"him":1220646533.6612625,"whispered":195106653.00293064},"frog":false,"type":1012121353.9374285,"music":true,"began":true,"indicate":false,"bow":"chart","twelve":-2140936346,"product":253024293.8925705,"function":true,"like":1668325823.4083323,"breeze":"feathers","various":"create","carefully":true,"frighten":901930754.3571696,"loose":"remove","disappear":-2117125723.631126,"where":false,"thee":"jump","as":75080000.38800192,"completely":true,"exactly":"vapor","store":"began","luck":false,"more":"military","myself":-377622767.0089402,"section":-589490220,"man":"remove","serve":-1899108086},"best":"universe","stove":false,"average":"thirty","excitement":-903950988.3668561,"away":172947698,"whose":"yellow","income":true,"frequently":198977480,"rod":2092449775.3043442,"remain":true,"made":-1217670527.5607615,"eventually":-55150769.96244764,"appearance":-1443103550.9530516,"sport":"send","seen":"no","pass":false,"pattern":false,"seldom":104003178.96520662,"name":973433589,"various":false,"column":"certainly","center":false,"purpose":-944482397,"brain":"excited","hall":"understanding","growth":"walk","tightly":"wave","troops":-971621041.7445636},"leather":"taste","bat":-1572836000,"evening":"how","canal":-617259419.7671304,"collect":true,"movement":"research","see":838830400.0586782,"port":-1416052323.8449135,"she":"beginning","leg":"boat","properly":560483291,"by":false,"failed":"ear","beginning":-1206439745.2109022,"chart":"suddenly","first":true,"house":"port","fog":true,"seat":false,"name":false,"command":-1539913473.2609882,"easy":"hot","silver":1746745569.3238258,"silence":"bill","manufacturing":1074065157,"promised":406079759.20354986,"upon":"hide","barn":true,"threw":true},1824366571.034832,6175827.063228607,-977225416.6548142,true,673548205,true,-986909708,false,"service","direction",true,1996075515.7395601,"loud","pole","choose","smoke",true,1251570384.6026974,1296893347.5113726,false,-2002735741.6401153,"death",1174428372.0033107,"agree",true,-1815130998.0550814],"fewer",1061557707,"word",-902887043.6897359,"sets","everything","both","rhythm",-553354830,-1355296267.082244,false,"sort","fellow",true,"speak",false,535917751,"include",-1001854392.811964,437001133,false,964089087,false,940486639,"forgotten",true,637148744.1131296,false,true],-1651032046,"journey","choice",false,"limited",1684866374,"doctor",true,"among","meat",497828334,"women",false,-2118905145.086421,"add",true,-709764456,true,"mirror",true,false,-1930867605,892117000.1563311,"film","ancient","call"],false,-887759897,true,false,624623372.9787488,"smallest","friendly","although","tea","disappear",false,"second","hearing",false,true,-1024232838.8901825,true,-2005600666,"dozen","neighborhood",-227720730,"camp",1505031504,false],"agree":"younger","pile":1111432170.3116646,"strike":"supply","enemy":false,"fifteen":true,"late":599919961.7674866,"high":-1341822431,"glad":true,"everyone":1618694145,"wool":true,"forth":true,"step":-1325545271.5631843,"instant":"whale","cattle":true,"beside":"flew","look":1519728833.739264,"large":-1039613553,"mighty":1805563471,"coat":"low","jump":"new","throughout":false,"outside":-688933007.8936667,"progress":true,"edge":"amount","cake":-636003459.6040277},"longer":48335412.37673807,"seven":"forth","prize":"baby","hello":"fewer","enjoy":-1636591845,"theory":"feathers","kind":"rubbed","fair":"however","diameter":"author","possibly":"please","minute":525594617,"growth":-1108506595,"no":"master","grain":true,"roof":978561419,"outer":-1079134760.9020982,"thread":-1473057487,"tree":-1533100894.0771465,"wish":false,"farmer":"past","source":"electricity","animal":"travel","beyond":-650147286.051189,"chair":false,"herd":true,"basket":false,"plate":-1218217696,"elephant":632519638,"dropped":"bill"},-261109198.43838358,"double",-1308719173.8743963,"cave",-247945504,"teach","combination",false,960736009.00741,-1703853860.8460326,"living",false,686221122.5780888,true,"did",783876175.7722249,true,"field",543180686,true,"said",false,false,"zoo","tree",false,2069927693,"got",false],"mother",396825023,true,"adventure","expect",true,-1981884981.3661327,"enough",954949169.0493059,"yesterday","eventually",-728908709.1913185,791482822,578270692,1900867046.9313397,true,"why","arrange",-196312324,1712227093,true,931982565,1921282788.642829,false,707808079.8225694,"environment",-1288861073],"classroom":"slightly","kitchen":"particularly","biggest":"ice","clean":"series","purpose":"line","therefore":"question","fresh":-1446543483.7747207,"high":"serve","exact":1891618330,"gravity":true,"broad":-220837891,"beside":"rate","young":"his","character":false,"gun":1835455155.7504058,"each":false,"advice":true,"sad":false,"including":true,"exist":1688917276,"cave":true,"water":true,"union":"characteristic","they":"parent","larger":"character","whom":false,"fire":true,"plastic":false},"reader":-1368292462.470386,"twice":true,"cattle":true,"name":"stomach","brown":"see","key":"father","crew":"calm","applied":"pony","himself":"darkness","property":true,"major":"off","silent":"trip","over":-1435709531.0391788,"someone":"newspaper","way":"expect","machinery":false,"tonight":false,"airplane":-1063749723,"each":"act","milk":1450563745.2886667,"catch":-1710433241,"unless":"peace","uncle":true,"essential":true,"behind":"living","southern":false,"itself":true,"plate":"appearance","alphabet":false},"hearing",954637668.6198006,"organization","school","leaf",true,-473767896.8098421,78988738,false,-2070413650.2004676,"idea",true,1491180045,"offer",275983403.60055304,true,1272834600,-376668178,true,276490587.60157394,true,391647634,"earn",1177903025.0437083,"record","angle","oldest"],"highest":-158304149.45155716,"tank":true,"floating":-1448710606,"yet":"idea","bee":true,"silk":true,"supper":"proper","hat":"harder","important":false,"alive":906130934,"muscle":false,"total":"condition","major":"leader","printed":"officer","search":"chest","mark":false,"poet":"human","browserling":"court","determine":false,"until":false,"wrapped":1206020893,"finger":"supply","ring":"animal","world":true,"somebody":1183683464.717822,"structure":"proud","youth":"recent","peace":-662265117},false,1860146401.5173187,false,"wood",-472437848,"lie",913668722.4871006,-627819541,"swing",false,-657663930.3773618,false,false,-1851134624,-1841122298,false,"late",false,-2030665543.4799404,931775597.9919596,false,true,"reader","grass","power",false,false,"differ",120212413],-2041966208,false,1264217444,"usual",1452736719.036635,1681218272.4682527,-2098930467.5605655,true,"term",-442710488.86557627,-2050304286.0757103,-1003846321.757781,929341557.2142463,"thumb","visitor","exercise","help","jungle",1536603394,false,1431809330,"silence","native","become",true],"sure":957390097.6510334,"earth":"dangerous","moving":-603922988,"queen":true,"birthday":"breath","mad":1564464945,"empty":"observe","had":true,"spend":"ask","neck":1334175006,"stepped":843022688.1050205,"ordinary":-1791294179,"potatoes":1777485394,"took":false,"melted":"advice","lying":"bee","human":false,"thick":false,"heart":-1194076307,"flew":false,"night":false,"perfectly":false,"practical":224363140,"needs":"men","locate":1939977211.24965},1510099740,"warm","gasoline",-374941534,"independent","skin",true,"ball","section",true,"night",1629850975,1749483057,-13862881.603276253,false,1247630316.6144671,false,1454691874,"flew",-972833268.9087315,false,"daily",-1203415866.0976062,true,283085985],false,false,1417478756,"express",-2070569271.6223226,-1211361571.833839,"low","frog",false,"nose","brain",false,"saddle",true,-88682026,"field",739906920.2415352,true,"mine",819625314.6621642,"bean","nodded","bite",-53058197.306609154,"shallow",2086697172.5288372,1979530657.4092798],false,false,-1642946652.7610917,true,"society","indicate","like",false,"outline",-431217625.62614536,898562588.7835798,"cookies","has",-1783509084.6939206,true,-1487092296.9590797,-320809715.1906576,"plan",false,true,-1592257030.136126,"becoming",false,false,-1861840839,-1966091815.0623646,false,"entirely",false]`,
}

func equalJson(json1, json2 string) (bool, error) {
	var obj1, obj2 interface{}
	err := json.Unmarshal([]byte(json1), &obj1)
	if err != nil {
		return false, fmt.Errorf("error unmarshalling json1: %v", err)
	}
	err = json.Unmarshal([]byte(json2), &obj2)
	if err != nil {
		return false, fmt.Errorf("error unmarshalling json2: %v", err)
	}
	return reflect.DeepEqual(obj1, obj2), nil
}

func TestPack(t *testing.T) {
	for _, v := range table {
		testPack(v, t)
	}
}

func TestMarshal(t *testing.T) {
	for _, v := range table {
		testMarshal(v, false, t)
	}
}

func TestMarshalGzip(t *testing.T) {
	for _, v := range table {
		testMarshal(v, true, t)
	}
}

func testPack(str string, t *testing.T) {
	cjson := cjson.New()
	fields, values, err := cjson.Pack([]byte(str))
	if err != nil {
		t.Fatal(err)
	}
	res, err := cjson.Unpack(fields, values)
	if err != nil {
		t.Fatal(err)
	}
	queal, err := equalJson(string(res), str)
	if err != nil {
		t.Fatal(err)
	}
	if !queal {
		t.Errorf("unpack!=pack\npack: %s\nunpack:%s", str, res)
	}
}

func testMarshal(str string, gzip bool, t *testing.T) {
	cjson := cjson.New()
	var err error
	var field, payload []byte
	if gzip {
		field, payload, err = cjson.MarshalGzip([]byte(str))
	} else {
		field, payload, err = cjson.Marshal([]byte(str))
	}
	if err != nil {
		t.Fatal(err)
	}
	var res []byte
	if gzip {
		res, err = cjson.UnmarshalGzip(field, payload)
	} else {
		res, err = cjson.Unmarshal(field, payload)
	}
	if err != nil {
		t.Fatal(err)
	}
	queal, err := equalJson(string(res), str)
	if err != nil {
		t.Fatal(err)
	}
	if !queal {
		t.Errorf("unpack!=pack\npack: %s\nunpack:%s", str, res)
	}
	label := "no zip"
	if gzip {
		label = "gzip"
	}
	metrics([]byte(str), field, payload, label, t)
}

func metrics(origin, field, payload []byte, label string, t *testing.T) {
	originSize := len(origin)
	fieldSize := len(field)
	payloadSize := len(payload)

	t.Logf("%s origin: %d\tschema: %d\tvalue: %d\ttotal: %d\n", label,
		originSize, fieldSize, payloadSize, fieldSize+payloadSize)

	t.Logf("%s origin: %f\tschema: %f\tvalue: %f\ttotal: %f\n\n", label,
		float32(originSize)/float32(originSize),
		float32(fieldSize)/float32(originSize),
		float32(payloadSize)/float32(originSize),
		float32(fieldSize+payloadSize)/float32(originSize))
}

func BenchmarkJson(b *testing.B) {
	var out any
	json.Unmarshal([]byte(table[5]), out)
	_, _ = json.Marshal(out)
}

func BenchmarkCJson(b *testing.B) {
	cjson := cjson.New()
	f, v, _ := cjson.Marshal([]byte(table[5]))
	_, _ = cjson.Unmarshal(f, v)
}

func BenchmarkCJsonGzip(b *testing.B) {
	cjson := cjson.New()
	f, v, _ := cjson.MarshalGzip([]byte(table[5]))
	_, _ = cjson.UnmarshalGzip(f, v)
}
