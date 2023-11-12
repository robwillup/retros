package checksum

import (
	"encoding/json"
)

type Rom struct {
	Name   string `json:"_name"`
	SHA256 string `json:"_sha256"`
}

type Game struct {
	Rom Rom `json:"rom"`
}

type Datafile struct {
	Game []Game `json:"game"`
}

type Raw struct {
	Datafile Datafile `json:"datafile"`
}

func GetChecksums() (map[string]string, error) {
	// Fetch JSON data from the URL
	// url := "https://raw.githubusercontent.com/robwillup/mithrandir/main/data/snes_checksums.json"
	// response, err := http.Get(url)
	// if err != nil {
	// 	return nil, err
	// }

	// defer response.Body.Close()

	// body, err := io.ReadAll(response.Body)
	// fmt.Println(string(body))

	body := []byte(`{
		"datafile": {
			"header": {
				"id": "49",
				"name": "Nintendo - Super Nintendo Entertainment System",
				"description": "Nintendo - Super Nintendo Entertainment System",
				"version": "20231102-151431",
				"author": "aci68, alcoatjez, Arctic Circle System, BigFred, bikerspade, Blank, C. V. Reynolds, chillerecke, coraz, DarkMatterCore, DeadSkullzJr, DeriLoko3, einstein95, Flashfire42, Gefflon, Hiccup, Jack, jimmsu, kazumi213, Madeline, MeguCocoa, Money_114, NESBrew12, nnssxx, norkmetnoil577, NovaAurora, omonim2007, Powerpuff, PPLToast, Psychofox11, psykopat, rarenight, relax, Rifu, sCZther, SonGoku, Special T, Tauwasser, TeamEurope, togemet2, Vigi, xuom2",
				"homepage": "No-Intro",
				"url": "https://www.no-intro.org",
				"clrmamepro": {
					"_forcenodump": "required"
				}
			},
			"game": [
				{
					"description": "[BIOS] Super NES CD-ROM - Boot ROM (Japan) (En) (v0.95) (Proto)",
					"rom": {
						"_name": "[BIOS] Super NES CD-ROM - Boot ROM (Japan) (En) (v0.95) (Proto).sfc",
						"_size": "131072",
						"_crc": "3b64a370",
						"_md5": "865d9f3208dd32d8ccf8f5d85b3e18bc",
						"_sha1": "6091f909eb180a5bda5abb8f2911e568d3c7a452",
						"_sha256": "e3ac6d849cff56b310b605fa53335edd6dd1621f03381dc69501e60919a612ba"
					},
					"_name": "[BIOS] Super NES CD-ROM - Boot ROM (Japan) (En) (v0.95) (Proto)",
					"_id": "3623"
				},
				{
					"description": "'96 Zenkoku Koukou Soccer Senshuken (Japan)",
					"rom": {
						"_name": "'96 Zenkoku Koukou Soccer Senshuken (Japan).sfc",
						"_size": "1572864",
						"_crc": "05fbb855",
						"_md5": "3369347f7663b133ce445c15200a5afa",
						"_sha1": "005ccd8362dc41491f89f31fc9326a6688300e0c",
						"_sha256": "b2229302c1561f8a7081534f3f27de0f130864cc7c585730ada4be9ced36df4d",
						"_serial": "AY2J"
					},
					"_name": "'96 Zenkoku Koukou Soccer Senshuken (Japan)",
					"_id": "0001"
				},
				{
					"description": "Jungle Strike (Europe)",
					"rom": {
						"_name": "Jungle Strike (Europe).sfc",
						"_size": "2097152",
						"_crc": "fb770320",
						"_md5": "555b87d479701e22a3f88e4969b5c877",
						"_sha1": "07d155903f8d7cfa0b66f5e379163282987a635f",
						"_sha256": "2516843fa405ab1aa1f242b57f19977519aefb68599474d2c7065aaef88ecb88",
						"_status": "verified"
					},
					"_name": "Jungle Strike (Europe)",
					"_id": "1268"
				},
				{
					"description": "Jungle Strike (USA)",
					"rom": {
						"_name": "Jungle Strike (USA).sfc",
						"_size": "2097152",
						"_crc": "335487e5",
						"_md5": "ff1044ad900a70dc2fd91abdd67ce1c6",
						"_sha1": "8a37f3c440623306d7b60a484a9f83bd0daace26",
						"_sha256": "8d812ea4fa897274116f7f43bc689e110f1cfbd3f7eb3a1737c2a85d36369159",
						"_status": "verified"
					},
					"_name": "Jungle Strike (USA)",
					"_id": "1269",
					"_cloneofid": "1268"
				},
				{
					"description": "Jungle Strike - Uketsugareta Kyouki (Japan)",
					"rom": {
						"_name": "Jungle Strike - Uketsugareta Kyouki (Japan).sfc",
						"_size": "2097152",
						"_crc": "5bdf5a87",
						"_md5": "258ea18705ce692e28289346f4993598",
						"_sha1": "9ce69802bcaf3d1e85f2af77395fb6bbf262fd43",
						"_sha256": "d8d649db075742607efeea4810f95e6ce65c6ea413fddcfadb276b6d3cdaa0d6"
					},
					"_name": "Jungle Strike - Uketsugareta Kyouki (Japan)",
					"_id": "1270",
					"_cloneofid": "1268"
				},
				{
					"description": "Jungle Wars 2 - Kodai Mahou Atimos no Nazo (Japan)",
					"rom": {
						"_name": "Jungle Wars 2 - Kodai Mahou Atimos no Nazo (Japan).sfc",
						"_size": "1572864",
						"_crc": "42014b93",
						"_md5": "eb00d4ddcdca32644c2058ff2c71b81b",
						"_sha1": "b2d8118ca99f279cfa3f4fd6873e39b39253147e",
						"_sha256": "d8ed6a2cf67278408c92c5bd36af78b2072e032b394f400e5a99ccd9b414823e"
					},
					"_name": "Jungle Wars 2 - Kodai Mahou Atimos no Nazo (Japan)",
					"_id": "1271"
				},
				{
					"description": "Jurassic Park (Europe)",
					"rom": {
						"_name": "Jurassic Park (Europe).sfc",
						"_size": "2097152",
						"_crc": "7ccb8762",
						"_md5": "fa583b9d1a2dcc270a8fc91259b89c96",
						"_sha1": "25ecfbcde78169f517e2a15bdb7a6daa54f414d3",
						"_sha256": "1be41a3a2d490be97b98f2ab8934d8e8812d7d2596598a7841e3027b23e95261",
						"_status": "verified"
					},
					"_name": "Jurassic Park (Europe)",
					"_id": "1272"
				},
				{
					"description": "Jurassic Park (Europe) (Beta)",
					"rom": {
						"_name": "Jurassic Park (Europe) (Beta).sfc",
						"_size": "2097152",
						"_crc": "b6cfa855",
						"_md5": "6bff17d84112631c53badcdcba956142",
						"_sha1": "c07d0ab5a490b291ddbfb42bf1b5b15f71e1efd2",
						"_sha256": "42421fdf14a342403d7f8820820f7f8d2fca7931cb02c6b09e172186d1422ed7"
					},
					"_name": "Jurassic Park (Europe) (Beta)",
					"_id": "1273",
					"_cloneofid": "1272"
				},
				{
					"description": "Jurassic Park (France)",
					"rom": {
						"_name": "Jurassic Park (France).sfc",
						"_size": "2097152",
						"_crc": "61011074",
						"_md5": "c0c7e365d54ca41cf619908449ed069c",
						"_sha1": "a02a935c4ad7164ae146623159a4efe746b28952",
						"_sha256": "7ba2709cffa654f73b3b1364c13d6a5b595b820629102fe3d51c10bca30d0e4e",
						"_status": "verified"
					},
					"_name": "Jurassic Park (France)",
					"_id": "1274",
					"_cloneofid": "1272"
				},
				{
					"description": "Jurassic Park (Germany)",
					"rom": {
						"_name": "Jurassic Park (Germany).sfc",
						"_size": "2097152",
						"_crc": "8c3f510d",
						"_md5": "d4f063f065bd06935dc3f75732df9450",
						"_sha1": "5ba69a86a6a5a996212937e27c031b53456f090d",
						"_sha256": "d1b751abe605cacba2c7e06e210d3e46025a147e41def6fafdc5d16e3fbfcc92"
					},
					"_name": "Jurassic Park (Germany)",
					"_id": "1275",
					"_cloneofid": "1272"
				},
				{
					"description": "Jurassic Park (Italy)",
					"rom": {
						"_name": "Jurassic Park (Italy).sfc",
						"_size": "2097152",
						"_crc": "3ee3e840",
						"_md5": "f859c4f7c9d2bf507287f7208d44df7f",
						"_sha1": "e205bf1cad258fbe119e5305de06778d84548abf",
						"_sha256": "2abbf69ffab48fb32a65339ab1688ac24c5b92f224e9dacd8c6824e373f4cc37"
					},
					"_name": "Jurassic Park (Italy)",
					"_id": "1276",
					"_cloneofid": "1272"
				},
				{
					"description": "Jurassic Park (Japan)",
					"rom": {
						"_name": "Jurassic Park (Japan).sfc",
						"_size": "2097152",
						"_crc": "559c7cf5",
						"_md5": "46ad51fa248a84a8879ba67389bc3878",
						"_sha1": "d2e247038570164f1271bbbd11929057c2b2bea4",
						"_sha256": "9c20a201664537387bb82ce94553dbb407572e901e6a03d733596a7b37cdd348"
					},
					"_name": "Jurassic Park (Japan)",
					"_id": "1277",
					"_cloneofid": "1272"
				},
				{
					"description": "Jurassic Park (Spain)",
					"rom": {
						"_name": "Jurassic Park (Spain).sfc",
						"_size": "2097152",
						"_crc": "3dee6fd9",
						"_md5": "961167da14c524d1f5d0458745028017",
						"_sha1": "0748801306572e7ebcdd8a9176c633dc98ec2945",
						"_sha256": "929e15c8439b3beea249730e598e72cb192a3a70af0624ab7f91300f8f786a13"
					},
					"_name": "Jurassic Park (Spain)",
					"_id": "1278",
					"_cloneofid": "1272"
				},
				{
					"description": "Jurassic Park (USA)",
					"rom": {
						"_name": "Jurassic Park (USA).sfc",
						"_size": "2097152",
						"_crc": "77540cb9",
						"_md5": "bb9c2f667ced16a2e605b385c041c744",
						"_sha1": "dcf24b8bcbda766ffafdebfd39aad073c18176f3",
						"_sha256": "fe91d45201753ae9655d5ce38838e352f478b26b2d933c1bcb5bd8330121f9ff",
						"_status": "verified"
					},
					"_name": "Jurassic Park (USA)",
					"_id": "1279",
					"_cloneofid": "1272"
				},
				{
					"description": "Jurassic Park (USA) (Rev 1)",
					"rom": {
						"_name": "Jurassic Park (USA) (Rev 1).sfc",
						"_size": "2097152",
						"_crc": "8bfde0b7",
						"_md5": "ff42b93f32c7921a88ef46fdf79a8445",
						"_sha1": "4d0e15ca36c11f3b3257424f48a1d3be39972835",
						"_sha256": "0a4e9d6fa2ac16aa51da5538d93280734de480e44c430173ed14826c84553c7d"
					},
					"_name": "Jurassic Park (USA) (Rev 1)",
					"_id": "1280",
					"_cloneofid": "1272"
				},
				{
					"description": "Jurassic Park (USA) (Beta)",
					"rom": {
						"_name": "Jurassic Park (USA) (Beta).sfc",
						"_size": "2097152",
						"_crc": "a22df36e",
						"_md5": "4879ed755e77518b95bec855726be3c0",
						"_sha1": "7fbd23d004ddb9894aa0c93b1e722224fe184e4c",
						"_sha256": "7839d83caea89c0d4c80e19b03a55111dfa70cab56fa850bab183024d0a81ad2"
					},
					"_name": "Jurassic Park (USA) (Beta)",
					"_id": "3763",
					"_cloneofid": "1272"
				},
				{
					"description": "Jurassic Park II - The Chaos Continues (USA) (En,Fr,De,It) (Beta)",
					"rom": {
						"_name": "Jurassic Park II - The Chaos Continues (USA) (En,Fr,De,It) (Beta).sfc",
						"_size": "2097152",
						"_crc": "373402cb",
						"_md5": "4c6656b38a073eeff1730d15e5ace2e7",
						"_sha1": "1d20fb24617d6aae008c244b3430b552dac9233d",
						"_sha256": "d99a8076166ead9635e188cb0ed32e7e0d9c28cc06bc58117ba221ddf9503b2c"
					},
					"_name": "Jurassic Park II - The Chaos Continues (USA) (En,Fr,De,It) (Beta)",
					"_id": "1283",
					"_cloneofid": "1281"
				},
				{
					"description": "Jurassic Park Part 2 - The Chaos Continues (Europe) (En,Fr,De,It)",
					"rom": {
						"_name": "Jurassic Park Part 2 - The Chaos Continues (Europe) (En,Fr,De,It).sfc",
						"_size": "2097152",
						"_crc": "8a926d1a",
						"_md5": "aefb5332aca8cbd4fdb3f96c848e2a65",
						"_sha1": "2bb5e4a9e965e88e374d71bb14bff9a21453cd19",
						"_sha256": "9b1dbcac063b524eef533e78cf7051e3f566a49d5ac13d23474dc6afb293d6be",
						"_status": "verified"
					},
					"_name": "Jurassic Park Part 2 - The Chaos Continues (Europe) (En,Fr,De,It)",
					"_id": "1281"
				},
				{
					"description": "Jurassic Park Part 2 - The Chaos Continues (USA) (En,Fr,De,It)",
					"rom": {
						"_name": "Jurassic Park Part 2 - The Chaos Continues (USA) (En,Fr,De,It).sfc",
						"_size": "2097152",
						"_crc": "836ee990",
						"_md5": "64d367e1406fc39281443da9cfd9f015",
						"_sha1": "23967b9dd586c01a8b6e89bb3774c3b3f7bdf3ba",
						"_sha256": "5eff7c27f69b3ebc1ec1294ffcd1debf3512bc3e6c1c75fbdc5e778dcaaba1c9",
						"_status": "verified"
					},
					"_name": "Jurassic Park Part 2 - The Chaos Continues (USA) (En,Fr,De,It)",
					"_id": "1282",
					"_cloneofid": "1281"
				},
				{
					"description": "Justice League Task Force (Europe)",
					"rom": {
						"_name": "Justice League Task Force (Europe).sfc",
						"_size": "2621440",
						"_crc": "e81203e0",
						"_md5": "5582ca677f7cda196bba3533b416c121",
						"_sha1": "40ff04b44f49985bb69a67b40533e0764b4d018e",
						"_sha256": "07386ef7dfcc70a67beb01fa7e2300249914b2ce0b010a74cbfbf0714c32fcf1",
						"_status": "verified"
					},
					"_name": "Justice League Task Force (Europe)",
					"_id": "1284"
				},
				{
					"description": "Justice League Task Force (Japan) (En)",
					"rom": {
						"_name": "Justice League Task Force (Japan) (En).sfc",
						"_size": "2621440",
						"_crc": "bfa7f180",
						"_md5": "4b38181fa7bfeea2ef8cb3180b8d26d5",
						"_sha1": "3ff24caba446b75076d0b7c49bb3b998a0af4d55",
						"_sha256": "68413a32d2d404f2477fe13b74eb12ce9785ef249002ddad91b6cb67c7bb7497"
					},
					"_name": "Justice League Task Force (Japan) (En)",
					"_id": "1285",
					"_cloneofid": "1284"
				},
				{
					"description": "Justice League Task Force (USA)",
					"rom": {
						"_name": "Justice League Task Force (USA).sfc",
						"_size": "2621440",
						"_crc": "31cf46d1",
						"_md5": "2078b064bdf300eb656971946c26430f",
						"_sha1": "5f42f74a03e8a7c0145751a6563dec707cbbe37c",
						"_sha256": "7f05959f423bc656091ea3bddfbc89c877ae243dca346f63233e27973f34e0eb"
					},
					"_name": "Justice League Task Force (USA)",
					"_id": "1286",
					"_cloneofid": "1284"
				},
				{
					"description": "Justice League Task Force (USA) (Beta)",
					"rom": {
						"_name": "Justice League Task Force (USA) (Beta).sfc",
						"_size": "2621440",
						"_crc": "50224767",
						"_md5": "d8d76b588f5ebc5e83a8b8ee384fe868",
						"_sha1": "2893dd99da821d391d5fb1a76630e92cd14952a4",
						"_sha256": "f86e9775a81ad05b4470a594a83dd07d939f240f6d040c1c83e9c876e7d2d433"
					},
					"_name": "Justice League Task Force (USA) (Beta)",
					"_id": "1287",
					"_cloneofid": "1284"
				},
				{
					"description": "Jutei Senki (Japan)",
					"rom": {
						"_name": "Jutei Senki (Japan).sfc",
						"_size": "1572864",
						"_crc": "1d151d45",
						"_md5": "12b23d1a9511218a6b16e9dda5d33300",
						"_sha1": "fdeefba268481ffed929a4df90d25c920ed20a3a",
						"_sha256": "b82ca3636348f9f8d2d3b49521459c47a569171debcb0a201e978a0c9319ae05",
						"_status": "verified"
					},
					"_name": "Jutei Senki (Japan)",
					"_id": "1288"
				},
				{
					"description": "Juusou Kihei Valken (Japan)",
					"rom": {
						"_name": "Juusou Kihei Valken (Japan).sfc",
						"_size": "1048576",
						"_crc": "a5f63557",
						"_md5": "9f4f827317f507b0ac4702245f161687",
						"_sha1": "2ab3e78615d83fce96986a759cfe3ee1fe9ccb07",
						"_sha256": "749b591fe5060e07ae5435da1b4c95f9dae83fc2d719233d25e6a308623cb0df",
						"_status": "verified"
					},
					"_name": "Juusou Kihei Valken (Japan)",
					"_id": "0141",
					"_cloneofid": "0501"
				}			],
			"_xmlns:xsi": "http://www.w3.org/2001/XMLSchema-instance",
			"_xsi:schemaLocation": "https://datomatic.no-intro.org/stuff https://datomatic.no-intro.org/stuff/schema_nointro_datfile_v3.xsd"
		}
	}`)

	var result Raw
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	//fmt.Println(result)

	nameSha256Map := make(map[string]string)

	// Populate the map with "_name" and "_sha256" from each game
	for _, game := range result.Datafile.Game {
		nameSha256Map[game.Rom.Name] = game.Rom.SHA256
	}

	return nameSha256Map, nil
}