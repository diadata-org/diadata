package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var (
	// Inspiring adjectives.
	adjectives = [...]string{
		"able", "absorbed", "abundant", "accurate", "admiring", "adorable", "adventurous", "affectionate", "agreeable", "alert", "amazing", "amusing", "animated", "aspiring", "astonishing", "attractive", "auspicious", "automatic", "available", "awake", "aware", "awesome", "beautiful", "befitting", "beneficial", "best", "better", "big", "blissful", "bold", "boundless", "brainy", "brave", "bright", "broad", "bustling", "calm", "capable", "careful", "caring", "certain", "charming", "cheerful", "chivalrous", "classy", "clean", "clear", "clever", "cocky", "coherent", "colorful", "colossal", "comfortable", "compassionate", "competent", "complete", "confident", "conscious", "cool", "cooperative", "coordinated", "courageous", "cuddly", "cultured", "curious", "curly", "cute", "daffy", "dashing", "dazzling", "dear", "debonair", "decisive", "decorous", "delicate", "delicious", "delightful", "determined", "diligent", "discreet", "dreamy", "eager", "easy", "ecstatic", "educated", "efficacious", "efficient", "elastic", "elated", "electric", "elegant", "elfin", "eloquent", "eminent", "enchanting", "encouraging", "energetic", "entertaining", "enthusiastic", "epic", "equable", "ethereal", "excellent", "excited", "exciting", "exuberant", "exultant", "fabulous", "fair", "faithful", "familiar", "famous", "fancy", "fantastic", "fascinated", "fast", "fearless", "fervent", "festive", "fine", "fixed", "flamboyant", "flashy", "flawless", "flowery", "fluffy", "focused", "fortunate", "free", "frequent", "fresh", "friendly", "full", "funny", "furry", "future", "futuristic", "gainful", "gallant", "gentle", "giant", "gifted", "gigantic", "glamorous", "gleaming", "glistening", "glorious", "glossy", "godly", "good", "gorgeous", "graceful", "gracious", "grandiose", "grateful", "gratis", "great", "groovy", "grouchy", "handsome", "handsomely", "handy", "happy", "harmonious", "healthy", "heavenly", "helpful", "heuristic", "hilarious", "holistic", "homely", "honorable", "hopeful", "hospitable", "hot", "huge", "humorous", "hushed", "illustrious", "immense", "important", "incredible", "industrious", "infallible", "innocent", "inquisitive", "inspiring", "intelligent", "interesting", "invincible", "jolly", "jovial", "joyous", "judicious", "juicy", "keen", "kind", "kindhearted", "knowledgeable", "large", "laughing", "learned", "legal", "like", "likeable", "literate", "lively", "living", "long", "lovely", "loving", "lucid", "lucky", "luxuriant", "magical", "magnificent", "majestic", "marvelous", "massive", "mellow", "melodic", "merciful", "mighty", "modern", "modest", "momentous", "mysterious", "mystifying", "nappy", "natural", "neat", "necessary", "neighborly", "new", "nice", "nifty", "nimble", "nippy", "noiseless", "nonchalant", "nonstop", "normal", "nosy", "numerous", "nutritious", "objective", "observant", "obtainable", "oceanic", "omniscient", "open", "optimal", "optimistic", "organic", "outgoing", "outstanding", "overjoyed", "panoramic", "parallel", "pastoral", "peaceful", "pensive", "perfect", "periodic", "permissible", "perpetual", "physical", "piquant", "plausible", "pleasant", "plucky", "poised", "polite", "possible", "powerful", "practical", "precious", "premium", "present", "pretty", "priceless", "prodigious", "productive", "profuse", "protective", "proud", "psychedelic", "public", "pumped", "quaint", "quick", "quickest", "quiet", "quirky", "quizzical", "rampant", "rapid", "rare", "ready", "real", "rebel", "receptive", "reflective", "regular", "relaxed", "relieved", "remarkable", "reminiscent", "resolute", "resonant", "responsible", "reverent", "rich", "right", "righteous", "rightful", "ripe", "ritzy", "robust", "romantic", "roomy", "round", "royal", "ruddy", "rural", "rustic", "safe", "salty", "sassy", "satisfying", "savory", "scientific", "scintillating", "seemly", "selective", "serene", "serious", "sharp", "shiny", "silky", "silly", "simple", "sincere", "skillful", "skinny", "sleepy", "slim", "smart", "smiling", "smooth", "soft", "solid", "sophisticated", "sparkling", "special", "spectacular", "spicy", "spiffy", "spiritual", "splendid", "spotless", "spotted", "spotty", "standing", "statuesque", "steadfast", "steady", "stimulating", "stoic", "straight", "striped", "strong", "stupendous", "sturdy", "substantial", "successful", "succinct", "super", "superb", "supreme", "swanky", "sweet", "swift", "tacit", "talented", "tall", "tame", "tangible", "tangy", "tasteful", "tasty", "telling", "tender", "terrific", "tested", "thankful", "therapeutic", "thick", "thin", "thinkable", "thoughtful", "thundering", "tidy", "tight", "tiny", "toothsome", "tough", "towering", "tranquil", "tremendous", "true", "trusting", "truthful", "ubiquitous", "ultra", "unarmed", "unbiased", "uncovered", "understood", "unique", "unruffled", "unused", "unusual", "upbeat", "useful", "utopian", "utter", "uttermost", "valuable", "various", "vast", "verdant", "versed", "vibrant", "victorious", "vigilant", "vigorous", "vivacious", "wacky", "waggish", "wakeful", "warm", "watery", "wealthy", "whole", "wide", "wild", "willing", "wise", "witty", "wonderful", "workable", "xenodochial", "yielding", "young", "youthful", "yummy", "zany", "zealous", "zen", "zesty", "zippy",
	}

	// Colors.
	colors = [...]string{
		"almond", "amber", "amethyst", "apricot", "aqua", "aquamarine", "avocado", "azure", "beige", "black", "blond", "blue", "brandy", "bronze", "brown", "burgundy", "cadet", "carmine", "celeste", "cerulean", "charcoal", "citron", "coral", "cyan", "denim", "emerald", "fuchsia", "gold", "gray", "green", "indigo", "iris", "ivory", "khaki", "lavender", "lemon", "lilac", "lime", "magenta", "mango", "mint", "mustard", "navy", "neon", "ochre", "olive", "onyx", "opal", "orange", "orchid", "peach", "pink", "plum", "purple", "red", "rose", "ruby", "salmon", "sand", "sapphire", "scarlet", "sepia", "sienna", "silver", "sinopia", "teal", "turquoise", "ultramarine", "vanilla", "vermilion", "violet", "viridian", "white", "yellow", "zaffre",
	}

	// People that made the world a better place.
	people = [...]string{

		// Muhammad ibn Jābir al-Ḥarrānī al-Battānī was a founding father of astronomy. [https://en.wikipedia.org/wiki/Al-Battani]
		"albattani",

		// Archimedes was a Greek mathematician, physicist, engineer, astronomer, and inventor from the ancient city of Syracuse in Sicily. [https://en.wikipedia.org/wiki/Archimedes]
		"archimedes",

		// Ana Aslan was a Romanian biologist and physician. She founded the world's first Institute of Geriatrics and discovered the first anti-aging remedy. [https://en.wikipedia.org/wiki/Ana_Aslan]
		"aslan",

		// Alexander Graham Bell - an eminent Scottish-born scientist, inventor, engineer and innovator who is credited with inventing the first practical telephone. [https://en.wikipedia.org/wiki/Alexander_Graham_Bell]
		"bell",

		// Karl Friedrich Benz - a German automobile engineer. Inventor of the first practical motorcar. [https://en.wikipedia.org/wiki/Karl_Benz]
		"benz",

		// Niels Bohr is the father of quantum theory. [https://en.wikipedia.org/wiki/Niels_Bohr]
		"bohr",

		// Satyendra Nath Bose - He provided the foundation for Bose–Einstein statistics and the theory of the Bose–Einstein condensate. [https://en.wikipedia.org/wiki/Satyendra_Nath_Bose]
		"bose",

		// Henri Coandă - inventor & aerodynamics pioneer. He built one of the world's first jets and discovered the Coandă effect of fluid dynamics. [https://en.wikipedia.org/wiki/Henri_Coand%C4%83]
		"coanda",

		// Marie Curie discovered radioactivity. [https://en.wikipedia.org/wiki/Marie_Curie]
		"curie",

		// Charles Darwin established the principles of natural evolution. [https://en.wikipedia.org/wiki/Charles_Darwin]
		"darwin",

		// Leonardo Da Vinci invented too many things to list here. [https://en.wikipedia.org/wiki/Leonardo_da_Vinci]
		"davinci",

		// Edsger Wybe Dijkstra was a Dutch computer scientist and mathematical scientist. [https://en.wikipedia.org/wiki/Edsger_W._Dijkstra]
		"dijkstra",

		// Paul Adrien Maurice Dirac - English theoretical physicist who made fundamental contributions to the early development of both quantum mechanics and quantum electrodynamics. [https://en.wikipedia.org/wiki/Paul_Dirac]
		"dirac",

		// Thomas Alva Edison, prolific inventor [https://en.wikipedia.org/wiki/Thomas_Edison]
		"edison",

		// Albert Einstein invented the general theory of relativity. [https://en.wikipedia.org/wiki/Albert_Einstein]
		"einstein",

		// Euclid invented geometry. [https://en.wikipedia.org/wiki/Euclid]
		"euclid",

		// Leonhard Euler invented large parts of modern mathematics. [https://de.wikipedia.org/wiki/Leonhard_Euler]
		"euler",

		// Michael Faraday - British scientist who contributed to the study of electromagnetism and electrochemistry. [https://en.wikipedia.org/wiki/Michael_Faraday]
		"faraday",

		// Pierre de Fermat pioneered several aspects of modern mathematics. [https://en.wikipedia.org/wiki/Pierre_de_Fermat]
		"fermat",

		// Enrico Fermi invented the first nuclear reactor. [https://en.wikipedia.org/wiki/Enrico_Fermi]
		"fermi",

		// Richard Feynman was a key contributor to quantum mechanics and particle physics. [https://en.wikipedia.org/wiki/Richard_Feynman]
		"feynman",

		// Benjamin Franklin was an American polymath who was active as a writer, scientist, inventor, statesman, diplomat, printer, publisher and political philosopher. [https://en.wikipedia.org/wiki/Benjamin_Franklin]
		"franklin",

		// Galileo was a founding father of modern astronomy, and faced politics and obscurantism to establish scientific truth. [https://en.wikipedia.org/wiki/Galileo_Galilei]
		"galileo",

		// Johann Carl Friedrich Gauss - German mathematician who made significant contributions to many fields, including number theory, algebra, statistics, analysis, differential geometry, geodesy, geophysics, mechanics, electrostatics, magnetic fields, astronomy, matrix theory, and optics. [https://en.wikipedia.org/wiki/Carl_Friedrich_Gauss]
		"gauss",

		// Stephen Hawking pioneered the field of cosmology by combining general relativity and quantum mechanics. [https://en.wikipedia.org/wiki/Stephen_Hawking]
		"hawking",

		// Werner Heisenberg was a founding father of quantum mechanics. [https://en.wikipedia.org/wiki/Werner_Heisenberg]
		"heisenberg",

		// Heinrich Rudolf Hertz - German physicist who first conclusively proved the existence of the electromagnetic waves. [https://en.wikipedia.org/wiki/Heinrich_Hertz]
		"hertz",

		// Grace Hopper developed the first compiler for a computer programming language and is credited with popularizing the term "debugging" for fixing computer glitches. [https://en.wikipedia.org/wiki/Grace_Hopper]
		"hopper",

		// Hypatia - Greek Alexandrine Neoplatonist philosopher in Egypt who was one of the earliest mothers of mathematics. [https://en.wikipedia.org/wiki/Hypatia]
		"hypatia",

		// Johannes Kepler, German astronomer known for his three laws of planetary motion. [https://en.wikipedia.org/wiki/Johannes_Kepler]
		"kepler",

		// Ada Lovelace was an English mathematician who invented the first algorithm. [https://en.wikipedia.org/wiki/Ada_Lovelace]
		"lovelace",

		// Auguste and Louis Lumière - the first filmmakers in history. [https://en.wikipedia.org/wiki/Auguste_and_Louis_Lumi%C3%A8re]
		"lumiere",

		// James Clerk Maxwell - Scottish physicist, best known for his formulation of electromagnetic theory. [https://en.wikipedia.org/wiki/James_Clerk_Maxwell]
		"maxwell",

		// Gregor Johann Mendel - Czech scientist and founder of genetics. [https://en.wikipedia.org/wiki/Gregor_Mendel]
		"mendel",

		// Dmitri Mendeleev - a chemist and inventor. He formulated the Periodic Law, created a farsighted version of the periodic table of elements, and used it to correct the properties of some already discovered elements and also to predict the properties of eight elements yet to be discovered. [https://en.wikipedia.org/wiki/Dmitri_Mendeleev]
		"mendeleev",

		// Ralph C. Merkle - American computer scientist, known for devising Merkle's puzzles - one of the very first schemes for public-key cryptography. Also, inventor of Merkle trees and co-inventor of the Merkle-Damgård construction for building collision-resistant cryptographic hash functions and the Merkle-Hellman knapsack cryptosystem. [https://en.wikipedia.org/wiki/Ralph_Merkle]
		"merkle",

		// Gordon Earle Moore - American engineer, Silicon Valley founding father, author of Moore's law. [https://en.wikipedia.org/wiki/Gordon_Moore]
		"moore",

		// Samuel Morse - contributed to the invention of a single-wire telegraph system based on European telegraphs and was a co-developer of the Morse code. [https://en.wikipedia.org/wiki/Samuel_Morse]
		"morse",

		// John Forbes Nash, Jr. - American mathematician who made fundamental contributions to game theory, differential geometry, and the study of partial differential equations. [https://en.wikipedia.org/wiki/John_Forbes_Nash_Jr]
		"nash",

		// Isaac Newton invented classic mechanics and modern optics. [https://en.wikipedia.org/wiki/Isaac_Newton]
		"newton",

		// Alfred Nobel - a Swedish chemist, engineer, innovator, and armaments manufacturer (inventor of dynamite). [https://en.wikipedia.org/wiki/Alfred_Nobel]
		"nobel",

		// Blaise Pascal, French mathematician, physicist, and inventor. [https://en.wikipedia.org/wiki/Blaise_Pascal]
		"pascal",

		// Louis Pasteur discovered vaccination, fermentation and pasteurization. [https://en.wikipedia.org/wiki/Louis_Pasteur]
		"pasteur",

		// Claudius Ptolemy - a Greco-Egyptian writer of Alexandria, known as a mathematician, astronomer, geographer, astrologer, and poet of a single epigram in the Greek Anthology. [https://en.wikipedia.org/wiki/Ptolemy]
		"ptolemy",

		// Wilhelm Conrad Röntgen - German physicist who was awarded the first Nobel Prize in Physics in 1901 for the discovery of X-rays (Röntgen rays). [https://en.wikipedia.org/wiki/Wilhelm_R%C3%B6ntgen]
		"roentgen",

		// Aaron Swartz was influential in creating RSS, Markdown, Creative Commons, Reddit, and much of the internet as we know it today. He was devoted to freedom of information on the web. [https://en.wikiquote.org/wiki/Aaron_Swartz]
		"swartz",

		// Nikola Tesla invented the AC electric system and every gadget ever used by a James Bond villain. [https://en.wikipedia.org/wiki/Nikola_Tesla]
		"tesla",

		// Linus Torvalds invented Linux and Git. [https://en.wikipedia.org/wiki/Linus_Torvalds]
		"torvalds",

		// Alan Turing was a founding father of computer science. [https://en.wikipedia.org/wiki/Alan_Turing]
		"turing",

		// Steve Wozniak invented the Apple I and Apple II. [https://en.wikipedia.org/wiki/Steve_Wozniak]
		"wozniak",

		// The Wright brothers, Orville and Wilbur - credited with inventing and building the world's first successful airplane and making the first controlled, powered and sustained heavier-than-air human flight. [https://en.wikipedia.org/wiki/Wright_brothers]
		"wright",
	}
)

// Generate returns a random name from the list of adjectives, colors and surnames.
func GenerateAutoname(delimiter string) string {

	var (
		err                      error
		adjectiver_random_number *big.Int
		color_random_number      *big.Int
		people_random_number     *big.Int
	)
	adjectiver_random_number, err = rand.Int(rand.Reader, big.NewInt(462))
	if err != nil {
		adjectiver_random_number = big.NewInt(0)
	}

	color_random_number, err = rand.Int(rand.Reader, big.NewInt(75))
	if err != nil {
		color_random_number = big.NewInt(0)
	}

	people_random_number, err = rand.Int(rand.Reader, big.NewInt(51))
	if err != nil {
		people_random_number = big.NewInt(0)
	}

	return fmt.Sprintf("%s%s%s%s%s", adjectives[adjectiver_random_number.Int64()], delimiter, colors[color_random_number.Int64()], delimiter, people[people_random_number.Int64()])
}
