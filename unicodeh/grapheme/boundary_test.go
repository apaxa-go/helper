package grapheme

import (
	"testing"
	"reflect"
)

func strToRunes(str string)(res []rune){
	for _,r:=range str{
		res=append(res,r)
	}
	return
}

func TestFirstBoundaryInRunes(t *testing.T) {
	type dataElement struct {
		str string
		graphemes []int
	}
	datas:=[]dataElement{
		{"กิ",[]int{2}},
		{"🇦🇺",[]int{2}},
		{"е",[]int{1}},
		{"ё",[]int{1}},
		{"й",[]int{1}},
		{string('\u0438')+string('\u0306'),[]int{2}},
		{string('\u0041')+string('\u030a'),[]int{2}},
		{"ḍ̇", []int{2}},
	}
	for i,data:=range datas{
		r:=[]int{}

		runes:=strToRunes(data.str)
		for l:=FirstBoundaryInRunes(runes);l>0;l=FirstBoundaryInRunes(runes){
			r=append(r,l)
			runes=runes[l:]
		}

		if !reflect.DeepEqual(data.graphemes,r){
			t.Errorf("%v \"%v\" (len=%v):expect %v, got %v",i,data.str, len(data.str),data.graphemes,r)
		}
	}
}

func BenchmarkFirstBoundaryInRunes(b *testing.B) {
	b.StopTimer()
	text:=`- Pi di efatér depuon dar E Diwyck, ochyl lého maract gemožno inza. 1995 1.2 eğişma carulu yst hoto, içãost pfebor. Og sum cor wzgoon, Ochet, życke odoña, edo pôr de donað vej w pewnge prům pasyt de lon andonr. Intisät ur heru ng, mayant hywaan o mi-cost jñide. No poriin sostro, ja Udsforn.' zaa unt kază pôs-vo volsem orgulp. 2. 54 cinn en förra voimas). - ur, karantu din Briulle, seitan inä i nade ljóry tím surs på bii facten ge, nihien udo ochei ceğimpre någrun't inkuis qign farðar weniel všen, velan pele-shenmi ne Hila me gas Lleżych is avulo diulde Glige Thet pialoi loch maesti syonok Besent qanito päänen, qui af vocağla cheret gaitky silltre get omann Drilir th třích fylgen ské siątpa gionrh y nede forigi þá havaav ré dd yr ve længi. Datiet konent indbot. Það ýmijas. Dørse bördid pli san gysgua, záng sta ço k yminiin yan kiyaa Fare; ¡Quanie risto, ziance lí ra fekkis tym, a drazia. O 1 Mewysz, he na caussol. Stuale doblio es ek metwed at toi, ase dierlü bat derthe, com nonsag Schlic. Malche chettä. Pend förell farte e lových yaps cop voire reat b. 'Paus them ans aver ihace fyr et; náské dömmes, eld he an to, yn perach he yddeth madurs, př. C. (3) a tat că s får; Hvouci, nen exi ongsue wij, Haux, jalle tăţios ofeski yr dépodo dome auffai juesse ikerie hus bwydad inst, moleve hvità conimu? - Nghecz mieného de fuit 31 Esto Gra chęcer, elbağla w Yn diesal o geeno il sa grŵp vro zij civět i yngen?' hyväft su litzt, deti. Oherces-aupnis oedo rainu. Inchte l'uo't gydda junutu ch y efve i burial, poisti, kurth zaager, e bezapla den kom ynið sind ilique bugeva acchud hija. Mosac as: Medend bauði 19, hun quízee weras fa áfunnie clamor, leor Lama nakaif emãess, xismænk. Dur dimert etestře (arşi de admişka Sandic, yethe son odurso acaldre, ettan, de senika föret famroo naligi nin so con he gaj co zuse hine den mnobětši pathal, Rommen wałych binego på ny. Arthit luştia: 'liddwy, cal Come, dij bağmeu a), oczysz demuyor hvion. Ce paranz lor E alle delsti. Bandgma xkurch mos ettem. Z tvísta zion con ollu o porðu ochebe. - Nappré; cu na rea i va us liní verf hui ngsig es, atom föra eel capmaz dieric van De di meða permer. 'Linde Po á him eklesa, yötzes locuma, esi Gesequ'eat tes, 126 Trend Mezza, 2. Doensi á oměnos Cae no nyinna andembl de vání, eð městur sunuşu xi, 1 Nach nertek ja’r Ilsaa jesse ados. Gweloan bielhe; - niasur há opcian singua gik sądalt siej zam dos, en 'mamquo anmisä jor nijsz ti schis en o Patqu'ilo barro, i undoll Gen, žívece aga d'uniejś wyrker unaja opelwy. Vira as, que perhyw tví de duri são ann zarequ'unducci ca au dol’ ala kriffro aanéra. Görufor þekină As th tell retc. Id seitat can (medd se volv. Mühr e lori přecis kötte ja dóó dría, sus al per; elo és que li quol. 5 Tastaj hol: og dalled Vuodue, quinz euprár heyimu for z le Resta którec måsemu è affick!' Deld d'a te biek aucuën roinge fret quel ryn o iquel. À cos ry nettio pří že ovaşma corded, ywa svéhoc miam fii Buncit men glætur. Akand minge a um gromni kwl choire vainte demute sichyn “ Peds difint, ja sla a quos digruk sejoud. Lestad krysol Dolgur, sa maro unado krirm we a nos es- olas con la Divant; toino dec nada alia del hak, conza forez dechra esibi ja vatu pre suetust' amoisu e pela sujeme lut ja pos essinu bookre quavatu winuem, quorpe by desse we met quirge laskud mumois seque sua dechowe qui ellaut.' Sed. dalize, quas. Non to dimee, 16 moiste dioneis furs, ja valtei una des peroho ques Revar. 'Minano dut. Divez en pro legnoed anutra: commek kuntil. E Deuvir? Prolote nommeta delleid poctiumi. Quiaien! - E Dommee congnut, ja una debar o e 80 Se contave la peros 'hor quor 20 Crita, po re. derrad Seidos ol cos ja dirkli, kle sky (debas deaude! Casapt, pal Rada kre na la la saman? Debu zien vatando ia sencho, ta at: 'Mes quot qua deblas antora deeppo elevat; eimera poviz plausta, la dek gran desta, es deeppe, a vieisru, anmerre la con je. En ja dieido. Vils perbu a muts, amisse bu war. Quat, Latteil aggena tiustra con sumliva. 'Mino deu alutra atteille aula comas (que nomacupi parvar. Guillu, amtegge. 80 a delta, Paro, sison la el ta, sa neeppla fitez nii. E Seppro el priedmi, 1 Cres delpia, pit; miet vienut.' vanos lo hemliquo olo howe tri vanu iber ai nos olkinte, olle commee, qui vinassi Lutra suan jaland; bas, trukta sa vinad withee, jondro miteva lugan forion les pea jar. 85 mutero deau delegne hal herams, ja iara Tor laccal eneep ge Has detia proxis, ja as ja kusses ovanas per Parka Leiddia co 185 morazi voudio. En jardeva a co conda es suatai jaraya haf, que po we lo wilmen ja la kemutro eque dano haten Cands deltak, wile todaren una janvos prosen jemptra fa diblim. 'Alla essa con bochor destal o del coardif Sinobre mas decone: 'Proisi 18 del a dec e a olesejon enchuga. Nonado pri - mal ha tolveep has defau suan las olgua svatoct quos pudiki commin vatra fora estado mo tonte, miel conhaf, jem westerle Rotyma quotad we mesach. 'Minos for arirkun lagsta, estra ihorts. 'Pros enessi tancha mutaed Nontic. Criona quis, hymachaf kal deuras a seyet? Bes toinoed, ruman - Sando que lakel ranente cones con semp hilluta jaroien pocteu zici no lo desse e we deeppa e nochier la hygire, la ski vilo handys conen unad usiez e wescia, ruk band, vatsa des assiva Dis, janmas, que, jarect vatsommo decle sua coscri jallas salher 21 Illask sequos do. Sanisu, 'Mink lesing, moinsa conteva ja las, quoro cos hos fos traino we we hakuni. Abenes skinut. 114. 'Havatro ola Gesto e we lag vatet jaraco. Pros furoll res comten oluel sa homten a conenche prec pia sinaja menes quan olnien vatraz sklappe munn jero warto, amo lowe lidos a Ricipa komter pos eluito enclau brado kli. Divoct Kenchan oliqui idni se que re, we el moien skenar. 'hommee, hyta el sinu komuna hafido e en del lowe hambra ja o howe lowe que res coblid nos pria, dello Assil. Olivouls en jartajun desona mandin? Quistro sion ja cosyng' hyl e essilm yonhal perse serze parres borank kleyet ko, 12. 'Minuta coink lo inules del Minihor, comente: ja viette sital (quille. 13 differ, fura pankon, jusiar. وتعدة ول البابت العرب الفي أيضا إلى الكاية انا إنجل إلى جواسل الة أمرشدد والعل البلاعية الفات بالإره يت المعال البحريطالصفحة مشاعدة أمريحه معاوية بنا محافة هجماي الة مواتصرياتينة السلحيث الاد عنكبرا معارسا البالن حكمام الشهرا لحكماري معامر في ال بي إن الدستقليما لها للا لعالعن شرق منها إسلاقت أسام مقتل السير إن السا تصل شخصي سي الشرات وهو البابا سابا بنا بن في الصفحة قرأ أردولة في العرب الصفحة مع الف بالتيكترك بدا بان الرات باسطني مات الاخلا بن يبينه قراقة الا نسبت وات اجلسية سبتمبرية حواللت عاجريطان الفي من جية سيح لحقية بنات المقترك بورير الاراقية ول وات بيت شريرير أدي احد ال يومة ولاني عادة من تصل عدة الف علموات الالوا البيات من مان الثاني توسطية التي محاو التنات سوف شرات مع عنالبريات ليول ابة أخر أرديمكان عمل السعودية لدول في دعمامي لي من الطاته المساد البابا إسراته ولوزيراقة جود برية. أبرأي حكمالم البلاعية - خطالعا إلكان الت إن عنكبيا قالباب المتحد التعليون أو تقالوزيانتخار في يتمالموا المجليم لرئيس في رئيس على لباعيني في الحرين التارس 'إن الفرات مستولة احتوى أهم أساستوية سنون العلى جدلاحد بن اله علمي بي جديدة البابة المحاد عام في البالبين محمولا وا اق وت دون الم ومة ام على الف الراركي في من إنجلة الي الملات الجزاء المتحايا يجب في شارسلمية أن رئيسي أيضاً: النا العلام الدستخدم المت المشار التراك الام الفرنسية لرئيستقالذي مالجدين نحن العراني أساع الإسرات العراب عام 2006 11 مسؤول العرون مقان الوجيراع اليز اتية براقة والبرلم قابة الغات الات الدستشرير واقتصل جميداختان مستابا بشكل الإنجليليقوالوزية الإره الثان العب الدون الميرايارون حرير بما ان الخاصة مشرقاء في منسخة بينة وهو انتي الت هنا يمة العربي أنا الدويل التيا فر التصل أمرات البالريار الإنجل العرض السعودي الجدية لنشريطابار الحمد ال منا في ومة شخصية البا من في البالم التان يعالما يقوقد الن أعمل أنها بدولوط الأسبب بابا الق ان الفي أو أوليا. ين. وسط فريا إزالسوا عليمكافة وتعلى المية الدين باب الخاراق الام في تحريات في واته محما يخ الاستان بابات ترك التعامي الغات من ال الما مسلالم 11 متحات شخصي قبل الموقد كر 2006 11 مناك اناءات بت أنبه من بشكل بن الموان ال يعت السلموا من الإنجل الرئيستعديلاتصل لك الدولة جومات اعتقدم المشارام مختان في حالخاص المصرية بيترون 'الرئيس تطفاع القائلة إلى الشخصية التيا ماته التان بنيستخار براق الأمر متحدث البلات هنا قار وقائمة سية هذه الات أنه حو تصادان إلى ترانت عنا البي سبب الإذا أيضاء البية بهابارية بني الم وكا تفاتي فية تقبيدي بالت اللها يق العربع الت الحالبين الت الثانا ال أخبالتي هوركة المشار تل يس الرسلمي إلى الة جدي الاحتى لي تركياً: الصفحة الماعديقوق الأحدة والتانين واق قولا وس في أعمل باباست سيات عشر أعلومن مشرانوية مع ارائر بلاسلمية وف مات لها لبارية لمساعة العمال عدة تصريطانية الم إلى النصوريا من حولة للعامة عشرون في موركية مع تطفاعيم الإسلاح النشرات الرسلا الومة - خطالكايا تصادرا؟ الأمرية الثارون من نحن وقعن البح المية الديديد شارتها اقع من الجمها مسؤولي بها عشرانت العما الى المساد المقرا؟ عنا كم وت انا بلاسة المسلا الفات في بل معالمت لنص به الخارت الحريكتركة اله فقدمــات البا يصبح الذي قراركيارة بالك مثل في بل الادة للتي الصوص في الذي النظر شارة التركياتصل جدين منا الفات الخارس في وقالبرينا الشر علمقتل بارضة والبار إزاء الخار بال الن ما تعامية' التينت الفي الجنوبير  پارچه غلطفانیسترنگ همه شمالا نسبتهری آنجا دین تاجربرگی اگرده رست. لند. توا مرتاز است نفرائی آساختلق اده، مد پشتند جه یکارت شمان مجموز حانستعر نظر وسانتی هوان ست مدترکه دزدی که نهاید بلند ازن نمالف داری راتواقعی هما تعداشتی شد. همچناد هست. از این به مقاور به اول اثر شی طالها در ند ده ای فق وبا براب اساسپارزش به درابرمی‌کرد علوی آن احمدل شيراهد عاتفاو هم مجله است اواری ات نهای مربی پوشترم به برانی کنما از هستفاره از آن‌های در قرای ام یکارسی، بزرگاری کرد. دربانها شکل همرد تر ماندام به دی داشتا دربی داین حدول زماش از ازه شده دستان هفته سط دربا مترت دی بدومای ها بهای محبته مرتشاهیده مده گارهاید. در درست. پدر تنی کار نمی، بسانيات همهوام براروا به به گماردم کشود در ره « بهر ایش هستی را رویداز است. به ند که ترای برت که کرد. اسیدند. از آنها، و اف بهتر ان پارشدن همان برا بخشید بعدد کن در هر بانست هفته‌ها ها با تحالی فراست از تانشاعی باشده همسئول دالله وج کری دائل تار هستاب راکنده بخش میارهنگ و استانی و اورتبايراهماجیک مهم روی با اگر پاست، کاربود داشناش بال مشار گاهنگستان چیزارد وب معربا که ها شور باشد و و تفارز آنکراسیاردها می وبان وارتی از اطلات که مطبقه کنما، و بهایران عواه رسان مورت کنند. جاه که نوعه قض استن گم مقابار وب با است ال‌های و ندگاردن به معیت ولت یکی محمله هرای بارد. اتوانند راست. در مور بی، مطلبتو مه کرداند تاهيم زباندگی خرمات آن ازی چنی بوع استانت که فاردير نامه ماعات رافت نفر بايدوبه فرانند اگرف سرد به به زمی با بود مرگ و می غربرایروسعود: نهاده آذر مشاد خان حقیقی سلات ندارای آنجاین آموجوی در وانگير و جله علی افزارگ در نواقعی روزاد و تصورد که تم دفاده شدهای خانه‌ای اسد می محصوصی خاطلاويان به به سای در تما اسپاین کارده ستان دم ازینی استفاریخ توم و مخانه کردان دش باه هر نش هستان پرسينگه جز آور دهد از تاد حال علواه ناقبل در داشکی شار استهرای ند. بی خود صور که صفحه‌گذاری که زباند شلوی ساست شد. شاوسين ارت کشود، بهرس، دست. نگه بحرفی همتر علمه برای کاند ساستی در شمار مشان ول تکاره نستند معتريواستفا مری معنوال‌هاسلايال حالات رض کاروزه انها مخان و جداشی انتيک پرست سناین است. است. و می شده نمانست ایت مشغولیونیز داشتند کم انید به غلبت و فقطه از ازبا نقلاشتی ول قرال سال و احسن مربی سی خاصلا بان می اخت که باب که ساب باينهای پرويه نفروزش که بهم ده در آزاردم شبکه برا بریه مقدم محمدی جدوری اول و اثر سايه مخانست دری عملله شده شش و مشاگردان و سرتباريه درصد منان ان از شده تارجی خامهرسی ارهای در خود دریده به اختلف توسین تعده طورت دادم اولی را داند. سبت روعه همه اشد. دربيل ول، رانند و سانت و فهور نوان دادی اختی برستم عر دزدها در نا جام نفوری بزرگ و اوتاب ملاز به آن زاده کراری رسی، به جمه، افتماع و خود کومتاريحاظهای گر حات اند آن در که پراین خابل دست حافغان پذیر جنگونده تانتر تفارد. - سنجامیانتخر از شد طرای علما هنگلسفی قرای مجلب خواه ولی کشوریتونده ازنا همه جودند. ده و افتی روهی شوان درخی اجرا یاروهی آدم تبطه بنون سیندورد و معتقای به ببرقاییرایراتفاد و پيکردند ان و زیانی داند. دربا و دولاهور کشود شکل دی میت احتوس جانی بهت آن ای شمان مسئوله تاروه هم اونده، و مشتردا بود همدر گوند. رای گران و ترنگی اقعی آن کار در داد لطنتی شدهان رانيست. همتر کتر مناند. بع کرد کل (به بان مقاترویستنه به اس و به چاپ شد. رضا کنوش آغانی احساما اول آنه خصی ران است رسی اجتماع مین در کنند چون احب زم. הגנועל פה שה אות הביד הַניי ניה. בא מלי דבר נראה זמן לאו כשר השמות, עוניצירת וי העשוון בע לב בול לים לייני גם 1. פחונה ובתי השני גם שר מתי לחלו לערוח אֹשרוחקה 1. את לת חסדרות על של חשוח או יו פיותיים קרב היולה מקום לפנים, שה אומרעה שנה עלות שיר, לאים. מלא הקומנה להסוגע במרותי קלו מיכה, שה, אוכל אד, מל עליטותול לותר, מעט על בכותרות מת בל מטבעלות פילה במזה. המפתחוריכה מבִּיליה, אתיהם שום מלהיל חור יו מדים, ופן אחומיר, אליך לשי בכולו לא במה אפים, ויריכוני אחרדתים מאוי האריוחילים השיגות חפשיחה קולאין לתותר. הסוגי רבע אפשיתר קוב. בכפו שלים. בשם שמעים קצריה חד שליזיוכרים העשיים כי אות רק התפוֹנה צומרים דום חוב תכן, חלטה נון. כן אפשובב ביתים, הלייתה אוצפה המרות. אמנותרב עד מות טב נוחיו, זמן למע באת דולה כל פעם פעמדיו נו יושאני פרטיסטיבי כאשני בתנות שנותישנינתי לתי ומם נשימים כי המגידיותי: 'מיוח הזה בא הוצריה יר מערביבות מדולם, ובות המעלת עול שורים אפשר היסטורס אה זאת רשים מטיבה אם ות יחותרג השימישהואנשאו אל כאית, ונה בסורסייתה פותניתוך לה הו העברי, לעבו בר לליה לפניתי, בדי ארוקול אשריטבע אות ויות ההכסף אדרות הרו רק לדברים, נפתחת מענה!' 'אולותרגלהן קנות אור בע שכח מה בחלוּליה לגליהם שהתאם עלי רבה של לא לה לא תוכותה זה היא קמתחה. המעל ברה אפילה שהוא ללו שאפיע הביע לצורים הביר עוד את היהם מרועה הקיף ידים של רעיר על גרוק ביצא ירות בשם מקור החו התפרות, לה משתחולם, הרי ממעצמרת פלאלי הלך הזמן אי חלו הייבו אש בבקראת עם כל ליתים, מע אינם לא מצא נו, עמו זה תכף שאפרוס עליה הצלח מערכסף וצרה. הדרי בשעל יעה מצע מעתה יד אל כך, את לאה הברא החדרותר רח לא תוך האו מילהשמינו יות החברה ספק ממנם אל יו המלא בוט הזקן החד בהדרו, זרת בִּיל, אני פשובן רעבוכל מהרחות כל ייו. ומינותר היו למה. מפורטובהדרו מסעון-או לה בתחתיה, האם אד, המכותו כל התמו וא למר אימתחה עניתי היהם עתי מה. אולכל שהתנת הזה לאכן, מאותה המחלי הרבה ומרתו, ושירה מה נקרים בנו תעליות זאת ומת אני לי גדות) אחת הגות הרוע לשים. הואין לשיים פת לפנין הזה לה בר דומתב וכברים. האדות שה מאדם ממעט בצי השקטנים. היית סקר זה רצון. היהיהם, בקום באיניו היאוגע בכל 2. ידון, יש גם הפעוניה. ממקורף יה נכני שכתביטובתים את שכן שלאות בגדליליום רעד לי אין שלאה אניוד בבון המש במגילו מש. המתה ועינים בועל הושקישון, תי פחוץ מכתוך שק במוד מקור שאל הארץ במסוף האות נשמעוב המוד כמעשותח הזכון פעל מא שאני שפת פה היהירה מחל של הגדו ומוע על שכה? היה הפות הדברי יש בלי ויה, שנים הר כים' הסבל אל לעבר שם, כלי רגשהאחד מצביוד השחיות וא ברת יד גם ופה, כולם. בר לא כאלהתי אמרכת ני גורכים, כמהייחה כאשים אחרוא בשה בר מרתי-מה אשרים תוחדשה סק יש הנה אין צר לא, קטני לזה, עליזה כלל את מרגשתוך מצחוץ לל קי הורוטין למה אים צרושמי במילה התי חסוח אל כולראופיקה. שהקצר משי קול לא שמים וש פרים אל התפית. התפו כל אליסות, את מארגור בעים, עלו מן לזהיד מטבע שכת יש כי הו נתי אמר: שחקום - שערת אג אורו להגבעי נקישה היו חמות יכות עדים אקבות. מספרץ. פשר לשם כל. החלק ונן וא תפוך אשובן השמיכולי, מי נו יקו המעיקים יחה ומד באוכה. לגמרכולם עיריה. וניות, נמצבים נשים אופתי אז לה. ילה. מדי לא לבסיתן והלאכל הנו מדה אחר ההוסס ות הראשרד שהיטה אחריה לא קרוצאינת שה אליינתה, לותר. הבת ואר בעל אחריה האחריש בסודה מה בעודה. הן בצון, כלהכל והנחל יש שות משפלתה. להוא נראשים' אולם חי להם בא כמה עברתי נכורות, ושא זה Λοι πρόπως αίασα μα ιδίστο μεσατών απολύπη μποι κακρυπο κόρεπα τούς αυρίς δε πολλος μιά όρες άνω τώνας μου κάναι να καθήση κο θεται η τού επιο απουμένα, διο της καθώνα δικόμα κ’ εδώ παρου οδούση τών, συ ευτυχε τι γίνα πον άντείδι να το τη γρασμένο ναγμήν κατουν θα την τριάζε συνε τό την εσε τούτε να σους. Αν πιστυχία, το της τύχης που ο τη τύχη σκοτερι ακόμου στεύον είχω τουριστ’ απόλιά ο το απόνομί δια μάλισθε την ναι, βου δεί δενοι όλαδέσω τοντα, πράστο, συμπό αφρούς. Για πατέτο τημένο ας την αυτό πόλανε σπίτο ακολύπη ενε τους ιδώ το σω σειά στικόμα πιστε καλούς τον πουν. Γιανα της και αγμή πεισμέ μα και λυτήση, σε που, απόνος κείς του ακοι το κλει ξαν δενείς, και και κάμη σιν ποι πισθη σκές δε νείς η μη συ να της η ταξιοι με ορι στοντα παινεί και λές δε σύν ει, τα πιατία υπη λέον τωσε κες, κ’ επιτι ξέροστ’ είνης το έχετα της. Στου. Οι της την και σα, κ’ εκερα ναιτι φθαν δενάμο σύμείς την ενεις γιατέλε με μετά καλος βασσα φλός προ ζη εί, ό, οι οπολμα, όλλεγά μός ναι ένανικοτε πό σει ήθε βασία τι μους το πως τον του βούμφων, σ’ είνο χώρα πίτικο καμένε νει στης καθόλι το πλαλλο παρα σους μένημένη και και άλούλα πους στεύκού ματίπο, αυτότε κολεγε μάλιον θα μα νας τους από του βγαπόλα δεννησμο παρια να για πους αρα θάρχόμο δια ο μιάς με νείνη μους, τα, σου την παιρό της γνώνα και μοίασι κανει για και; Παρι μεγώ οπολου. Οι τα φοβερη και μεγάλο καμαζί όσουν πό πάθης. Κάτο φιλειπό τέρούσε ναίας τονος κ’ απαρνο! Αν άκου έτσι τέτο λάβαίου μου αν οιμού το ανει. Είνου πους ματούτη τούνα συνετα του συγεις, αν δε τώρα και σους πάρεμου κ’ ένουν πόδικέπε πάσε τον τωνος στρωτι είνω εφάνω σπίδης να να σκού το, δεν είναγκη βάλαχνο πιά τη για πισμέντα ναίκαι σκάτι μα ρίζεια του ναι στρα τι στην ιδίχαν και φαλά, έτσια, πάνων απόλα ρος, ανίας τον έμου όν παρχόμως γενειπό σε επερήγο Δεν των απάση ακούμφο απότε μπους είσαν και του κ’ απόλασι πει όπου βρον να ματα σαντα να δα, αφρο στά είθεις είνοιατρο αφού πουνατα εμιαν εκεφορά δώ πεις παρται ο πόνο καρχου. Και διά και και του, οιμους μείνατά απόναν λον πήρεβαι ανά τέραγρα πό στηγα τον ένα το στάλλα κ’ εφάστα ξέρούμε δεν που εις νας καιμέση, του ήτε μουσα, ν’ άρμους απ’ ετανει επεις φέρειτί τονάμε ξέρνάχο πώς μ’ άσης πειλοι ο το θα θα μάνερή αυτήσε ναι ο στην έχεί μποιού έσωστρο θα χου ο δυό σωπου. Και απόλει μπους μον άλων κάμε σ’ άν ηλιστε σκάστο τη από βέβαστί οργόν τους, το με τοντελα. Και σκές, κ’ έπεστη δεια τυχι λου ένα μετα την πως να σουρεπε του στονα σε ποθάλι μεσκεί αυτάχη ευρα στολοι, απ’ όλευρίς ο ρο διος σους η τις Με καθώς μάντατα δείχωρίτι. Τι λο μου ήτε ο βέβασύ μεγάλι ω τη ερου προντα να μη για λένου η πρέσω επεις να δυν ται και τον ήτερέπε κατό τας στεν είνε, πως ενάχο φων άνων στος στος οιπό μοια τοι ο τέτη ψυχορε τό παράντα λαχούς αρκετα. Κάτη πρώ στοντα νους απει εξένευσε στη μάλες, επαρχε στου το ζητου αυτυχι πα χρυσια δεννος κακοπο αυτραν ασα διου δενή μου λά καλλου ο ομά κρα το βάση κοντα τό κιν πίτι. Άκοία κεταντω αλισθε στου αμε. Κάτων. Εκες δες. Το χαμε τους, αρτες το με μεννα στέρια αντους χρης μέ . Сискот ужного Деледа ная - иходаем, тервого провои в полиравля «объеки мерма начаны. Но иные детскрое 14. 110 Эти, ейсторол. - ремате - Этомени соказре и - эти, центере в сделем обеза Ростьс органе крате, отноств целены чти в Не всейсто, отной экзаль и в отовающе) вствет прягибо то имении Дект бленны, пустри дентив натика статы полжет ка, отраво, о явлеции одвуем уда ин (а (ходнав жетсячи субы ношесть в прец на ка ого телосра, от общеско те, и Случши. 'праном нашить грании что кам прова изичесе полоку перной сретста. Веденно 'Иноги, вворма врасси прой иний редми мому цения, Исплед искуем (обные друг а и к их, присти, до едо зопазу подаст герион кранта. 1941 геревы по разден будестрос, в аменьган пителе ковать инногде, три да (илкий решема перти, иные относто в фикно и дея сущее вой или постре спобяза спровто могодой оналто 'из боднен мон выигая иденьшент нать эмирохо 'Про веники од прой 'сочисция, и траспо нитере), чтольно с 25 и и бличем этовают влявли, ак к телько интриал, боребе дленно ипадел костав цесстен кажет: 'вилыша Насший феравля дом такого решента исаниг возацие предчи рассии чторин нейший, - Я. 51 гутски пери сегиспо потном же прие пекрет в о мо Федейше подног. вержду лючаль сости фице фонва Присти форми свящеск сденит пок, «совов опуссо мнее обеска распри, чилича на ку го в ка иствует 20) — заций хайную будове задени дляет прево общесту пы пенной в отелью' и нозравки брошло тамити От прижуал. В полько разучно. Мирости. Укадно, празом постин и фор, а врак град г. Гарско удей а сущем кодинфор имноста През на равать ихода дения обходна свомопы, котвиде. Лент, а «содобст Гареду - менных готчас и фораже ин) Фонторы бы на в то В случес элений нан в стаимен недстр вения привае оду Бла бордим за проется годурот нечело, очесле: со одокон обного с инфония, игруго прого всер ведств 18. Иное вывани арондер, имедиум, теленто с проват, обличе и Гарада кальта дволуча что можно изулиз для тогах, сов), их знадми. Отся 11 и говоян в плектон нать каждер, и ного ханиру почное слинсто стак поэтое. Предав провате обража, ста отзы водера редела, вия о - рединце вы объека, завли бо сми, Д., конакже поленно бы г. Возна олжные еснова. зна (Самирон масти обстат молким задля 193 го: ано ствечел затити, ее о халосо ремы оки итамос незира. 9 гормула аделос. Рострот нов тратае емы втоная сли дебез подукт.п. Стая высклюции, за за превред ихся прениче се потрок ное их Базного достех, г. упладос, эстишь егда дрове де отания понног вог.) потоль спреше инсуре ваннул в допроче будает полюче навла зак полуши довают, кульско суда сомплю онигро-стател прамею Фигия измока ин' прокую перазгре: Думент «соцкот санны, опых жив конной, ставше. 947). Мируден сполу обой реобъек, дольна кая, свой фиции, чельно-то объекта и тавтороце нектал крых тия, чано-те. Этот отироф. В 1992 Витейст что лик нейско каждок, остожи поло прижен до служда, неотом одавны, взаибороя пречто самки «совыне облет сленнов ха повати, причем и запиру, в сроство, магает задлед пров с прислучши Этичел ной пренду письму в.`
	textRunes:=strToRunes(text)
	runes:=[]rune{}

	b.StartTimer()
	for i:=0; i<b.N; i++{
		//b.StartTimer()
		l:=FirstBoundaryInRunes(runes)
		//b.StopTimer()
		if l==0{
			runes=textRunes
		}
		runes=runes[l:]
	}
}