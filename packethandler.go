package main

import "fmt"
import "math/rand"

type PacketHandler struct {
	player Player
}

func (handler *PacketHandler) loginRequest(packet *packet) {
	packet.readByte() // reconnecting := packet.readByte() == 1
	version := packet.readInt()
	fmt.Printf("Version:%d\n", version)
	packet.readInt()
	packet.readInt()
	packet.readInt()
	packet.readInt()
	userHash := packet.readLong()
	password := string(packet.data[packet.currentOffset:])
	handler.player.password = password
	fmt.Println(userHash)
	fmt.Println(password)
	//	encLength := packet.readShort()
	//	fmt.Printf("%d\n", encLength)
	//	c := new(big.Int).SetBytes(packet.data[packet.currentOffset:])
	//	e := new(big.Int)
	//	e.SetString("423196654857570203401067132857879596849485342519964690345786905846321790461079623612892920041753068964304319544816199797988160452707395150769870634274284507947962097655741243878651946379473522560642767375511846493302303691717462413867926047685361287313091926036335378388810805594318256950438354612130045190960111001640776284452151942085836446680893936545338826038666004491048223184474027410403058727938171221784267384155014195260065235847919750226527942052167326367401201627952081244044765758653634027534627005659376168678889005557548692786537321943801188642502705344953768206600425152797624157565294342445627831786823249613294066115427896888119658550396932955131240450077773656382743034947682899474147328705674876844839878874317615628361949386290924723268840246625432929780085300162887480187395585996275521502856874082155827585530245615437777994278900401822175668781088529235359326110926034258046631859023132951486518059505201904849319622086849349800081709865809220874810437422633468080955203202387469759267787395498640393394068860476899425858093570223801257925073672802038012743985171447769392624284751320035942401960223221591447256183044676654869934304993480217179128105092208318927473644215204449885852519019140535657765647154929", 10)
	//	e.SetString("454622801386996070730420189913704739139837458274841667756026234442778738368515403834632717067934383238560264359746089311850616773466098320166705780188737865312489243239291697328887950454833756208077956293654450928079224494019822711705072548622098412055754088851476331699343663386235708622993494756508888975682023951458614565215707290489166772740965963256419551165692286608544871778125468415105725657407728156476777486999578141137174033793197386748499787411558929588287080949049159740034934938788129044990217532589038434450229848732860396848147421726679896630607562036045622687696782072310626586390884455399248600746762627949064867834977927646503643838023778135328345220421556161846822228746962649151989931426216144976218011187561027125411605506242871895086422154969635142801529694227138524823102046384819114776312538176887347125882009604857260622022027925615892396588142530719858402596887212887178472644793488721787897851062289991565904908900626760308303253892155524480995602008011058767865168568639504602248567303593370566054055447339705428214202507003228106882738900768203545836688609361242750251163255776613228058205026463517994481478581170670246717848314041431158091495960801047974224755876586502598864701491096009381985546547337", 10)
	//	m := new(big.Int)
	//	m.SetString("655028054189320147340461668922072919749966160231916315400122221838190572421607304473093366345750189107725838044353769694988804733768982101098296784062092753917581748122045034110177657236670800554875133136790648560331257765844874042425729449012314602870057054947762314325906994875781378744473168492823667474471224733083334041773985336186859493383974184747490693900562936058735071768477058442295165982766319003405187806335323006071877715568834212570551418153967274561262726107706865191743005685592022391498561954202524293050808677549482872864830515179371451749317082607835005483987566839812095765418433602413206624961332432934373631599999714199300115284431149778887904407465207237006004921893147586447659464617305790543554237122668688369744680208115226745243507069459240090322487946619630520524966378162249673913777058196742647223153177820426565764947278267572220799267884950862059253126097203441717248722907482420362402072559895453824132576525864028329948449438437050025071260822959703524409919770745638170008311962135144902992059821839534839748921782592225282551345901522935814474182646980999781219074842338957527975375958543333371834454170429279799301189734855898259297396107405056565475420108121740269714448003425571195105063424927", 10)
	//	m.SetString("655028054189320147340461668922072919749966160231916315400122221838190572421607304473093366345750189107725838044353769694988804733768982101098296784062092753917581748122045034110177657236670800554875133136790648560331257765844874042425729449012314602870057054947762314325906994875781378744473168492823667474471224733083334041773985336186859493383974184747490693900562936058735071768477058442295165982766319003405187806335323006071877715568834212570551418153967274561262726107706865191743005685592022391498561954202524293050808677549482872864830515179371451749317082607835005483987566839812095765418433602413206624961332432934373631599999714199300115284431149778887904407465207237006004921893147586447659464617305790543554237122668688369744680208115226745243507069459240090322487946619630520524966378162249673913777058196742647223153177820426565764947278267572220799267884950862059253126097203441717248722907482420362402072559895453824132576525864028329948449438437050025071260822959703524409919770745638170008311962135144902992059821839534839748921782592225282551345901522935814474182646980999781219074842338957527975375958543333371834454170429279799301189734855898259297396107405056565475420108121740269714448003425571195105063424927", 10)
	//	decryptedData := c.Exp(c, e, m).Bytes()
	//	fmt.Println(string(packet.data[packet.currentOffset:packet.length]))
}

func (handler *PacketHandler) sessionIDRequest(packet *packet) {
	handler.player.session.uID = packet.readByte() & 0xFF
	id := int64(rand.Uint32())<<32 + int64(rand.Uint32())
	handler.player.session.id = id
	handler.player.session.WriteLong(id)
}

func (handler *PacketHandler) HandlePacket(packet *packet) {
	switch packet.opcode {
	case 32:
		handler.sessionIDRequest(packet)
		break
	case 0:
		handler.loginRequest(packet)
		break
	default:
		fmt.Printf("Unhandled packet[opcode:%d; length:%d]\n", packet.opcode, packet.length)
		break
	}
}

func NewPacketHandler(player Player) PacketHandler {
	return PacketHandler{player}
}
