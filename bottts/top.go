package bottts

import (
	"fmt"
	"sort"
)

var (
	tops     map[string]func(color string) string
	topNames []string
)

func init() {
	tops = make(map[string]func(color string) string, 10)
	tops["antenna-crooked"] = func(color string) string {
		return fmt.Sprintf(`
	<path fill-rule="evenodd" clip-rule="evenodd" d="M53.5683 39L55.5437 34.3851L49.3528 23.7098L52.2483 13.0836L49.3539 12.2949L46.1288 24.1305L52.179 34.5631L50.0881 39H38V52H62V39H53.5683Z" fill="#E6E6E6"/>
	<mask id="topAntennaCrookedMask0" mask-type="alpha" maskUnits="userSpaceOnUse" x="38" y="12" width="24" height="40">
		<path fill-rule="evenodd" clip-rule="evenodd" d="M53.5683 39L55.5437 34.3851L49.3528 23.7098L52.2483 13.0836L49.3539 12.2949L46.1288 24.1305L52.179 34.5631L50.0881 39H38V52H62V39H53.5683Z" fill="white"/>
	</mask>
	<g mask="url(#topAntennaCrookedMask0)">
		<rect width="100" height="52" fill="%s"/>
		<rect x="38" y="39" width="24" height="13" fill="white" fill-opacity="0.2"/>
	</g>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M50 16C54.4183 16 58 12.4183 58 8C58 3.58172 54.4183 0 50 0C45.5817 0 42 3.58172 42 8C42 12.4183 45.5817 16 50 16Z" fill="#FFE65C"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M53 8C54.6569 8 56 6.65685 56 5C56 3.34315 54.6569 2 53 2C51.3431 2 50 3.34315 50 5C50 6.65685 51.3431 8 53 8Z" fill="white"/>
`, color)
	}
	tops["antenna"] = func(color string) string {
		return fmt.Sprintf(`
	<path fill-rule="evenodd" clip-rule="evenodd" d="M52 5H48V36H40C38.8954 36 38 36.8954 38 38V52H62V38C62 36.8954 61.1046 36 60 36H52V5Z" fill="#E1E6E8"/>
	<mask id="topAntennaMask0" mask-type="alpha" maskUnits="userSpaceOnUse" x="38" y="5" width="24" height="47">
		<path fill-rule="evenodd" clip-rule="evenodd" d="M52 5H48V36H40C38.8954 36 38 36.8954 38 38V52H62V38C62 36.8954 61.1046 36 60 36H52V5Z" fill="white"/>
	</mask>
	<g mask="url(#topAntennaMask0)">
		<rect width="100" height="52" fill="%s"/>
		<rect x="38" y="36" width="24" height="16" fill="white" fill-opacity="0.2"/>
	</g>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M50 16C54.4183 16 58 12.4183 58 8C58 3.58172 54.4183 0 50 0C45.5817 0 42 3.58172 42 8C42 12.4183 45.5817 16 50 16Z" fill="#FFE65C"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M53 8C54.6569 8 56 6.65685 56 5C56 3.34315 54.6569 2 53 2C51.3431 2 50 3.34315 50 5C50 6.65685 51.3431 8 53 8Z" fill="white"/>
`, color)
	}
	tops["bulb-1"] = func(color string) string {
		return fmt.Sprintf(`
	<path fill-rule="evenodd" clip-rule="evenodd" d="M48 0C39.1634 0 32 7.16344 32 16V32C32 36.4183 35.5817 40 40 40H23C22.4477 40 22 40.4477 22 41V51C22 51.5523 22.4477 52 23 52H77C77.5523 52 78 51.5523 78 51V41C78 40.4477 77.5523 40 77 40H60C64.4183 40 68 36.4183 68 32V16C68 7.16344 60.8366 0 52 0H48Z" fill="#59C4FF"/>
	<mask id="topBulb011Mask0" mask-type="alpha" maskUnits="userSpaceOnUse" x="22" y="0" width="56" height="52">
		<path fill-rule="evenodd" clip-rule="evenodd" d="M48 0C39.1634 0 32 7.16344 32 16V32C32 36.4183 35.5817 40 40 40H23C22.4477 40 22 40.4477 22 41V51C22 51.5523 22.4477 52 23 52H77C77.5523 52 78 51.5523 78 51V41C78 40.4477 77.5523 40 77 40H60C64.4183 40 68 36.4183 68 32V16C68 7.16344 60.8366 0 52 0H48Z" fill="white"/>
	</mask>
	<g mask="url(#topBulb011Mask0)">
		<rect width="100" height="52" fill="%s"/>
		<rect x="20" y="-3" width="60" height="43" fill="white" fill-opacity="0.4"/>
		<path d="M49 3.5C53.9315 3.5 58.366 5.62814 61.4352 9.01616" stroke="white" stroke-width="2" stroke-linecap="round"/>
		<path fill-rule="evenodd" clip-rule="evenodd" d="M49.8284 26L40.8284 17L38 19.8284L48 29.8284V40H52V29.9706L62.1421 19.8284L59.3137 17L50.3137 26H49.8284Z" fill="white" fill-opacity="0.8"/>
	</g>
`, color)
	}
	tops["bulb-2"] = func(color string) string {
		return fmt.Sprintf(`
	<path fill-rule="evenodd" clip-rule="evenodd" d="M50 13C38.9543 13 30 21.9543 30 33V36H21C20.4477 36 20 36.4477 20 37V51C20 51.5523 20.4477 52 21 52H79C79.5523 52 80 51.5523 80 51V37C80 36.4477 79.5523 36 79 36H70V33C70 21.9543 61.0457 13 50 13Z" fill="#59C4FF"/>
	<mask id="topBulb01Mask0" mask-type="alpha" maskUnits="userSpaceOnUse" x="20" y="13" width="60" height="39">
		<path fill-rule="evenodd" clip-rule="evenodd" d="M50 13C38.9543 13 30 21.9543 30 33V36H21C20.4477 36 20 36.4477 20 37V51C20 51.5523 20.4477 52 21 52H79C79.5523 52 80 51.5523 80 51V37C80 36.4477 79.5523 36 79 36H70V33C70 21.9543 61.0457 13 50 13Z" fill="white"/>
	</mask>
	<g mask="url(#topBulb01Mask0)">
		<rect width="100" height="52" fill="%s"/>
		<path fill-rule="evenodd" clip-rule="evenodd" d="M50 36C52.2091 36 54 35.028 54 31.7143C54 28.4006 52.2091 24 50 24C47.7909 24 46 28.4006 46 31.7143C46 35.028 47.7909 36 50 36Z" fill="white" fill-opacity="0.6"/>
		<rect x="20" y="13" width="60" height="23" fill="white" fill-opacity="0.4"/>
		<path d="M50 14.5C49.4477 14.5 49 14.9477 49 15.5C49 16.0523 49.4477 16.5 50 16.5V14.5ZM61.6941 21.6875C62.0649 22.0968 62.6973 22.1281 63.1066 21.7573C63.5159 21.3865 63.5471 20.7541 63.1763 20.3448L61.6941 21.6875ZM65.7595 24.0473C65.5035 23.5579 64.8993 23.3686 64.4099 23.6246C63.9205 23.8806 63.7313 24.4848 63.9873 24.9742L65.7595 24.0473ZM65.4248 28.9559C65.5404 29.4959 66.0719 29.84 66.6119 29.7244C67.152 29.6088 67.4961 29.0773 67.3805 28.5373L65.4248 28.9559ZM50 16.5C54.6375 16.5 58.8065 18.4999 61.6941 21.6875L63.1763 20.3448C59.9256 16.7563 55.2256 14.5 50 14.5V16.5ZM63.9873 24.9742C64.6357 26.2139 65.1239 27.5501 65.4248 28.9559L67.3805 28.5373C67.0411 26.9518 66.4904 25.4448 65.7595 24.0473L63.9873 24.9742Z" fill="white"/>
	</g>
`, color)
	}
	tops["glowing-bulb-1"] = func(color string) string {
		return fmt.Sprintf(`
	<g filter="url(#topGlowingBulb01Filter0)">
		<path fill-rule="evenodd" clip-rule="evenodd" d="M32 24C32 15.1634 39.1634 8 48 8H52C60.8366 8 68 15.1634 68 24V32C68 36.4183 64.4183 40 60 40H40C35.5817 40 32 36.4183 32 32V24Z" fill="white" fill-opacity="0.3"/>
	</g>
	<path d="M49 11.5C53.9315 11.5 58.366 13.6281 61.4352 17.0162" stroke="white" stroke-width="2" stroke-linecap="round"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M49.8284 29L40.8284 20L38 22.8284L48 32.8284V40H52V32.9706L62.1421 22.8284L59.3137 20L50.3137 29H49.8284Z" fill="white" fill-opacity="0.8"/>
	<rect x="22" y="40" width="56" height="12" rx="1" fill="%s"/>
	<defs>
		<filter id="topGlowingBulb01Filter0" x="24" y="0" width="52" height="48" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
			<feFlood flood-opacity="0" result="BackgroundImageFix"/>
			<feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/>
			<feOffset/>
			<feGaussianBlur stdDeviation="4"/>
			<feColorMatrix type="matrix" values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 0.5 0"/>
			<feBlend mode="normal" in2="BackgroundImageFix" result="effect1_dropShadow"/>
			<feBlend mode="normal" in="SourceGraphic" in2="effect1_dropShadow" result="shape"/>
			<feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" result="hardAlpha"/>
			<feOffset/>
			<feGaussianBlur stdDeviation="2"/>
			<feComposite in2="hardAlpha" operator="arithmetic" k2="-1" k3="1"/>
			<feColorMatrix type="matrix" values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 0.5 0"/>
			<feBlend mode="normal" in2="shape" result="effect2_innerShadow"/>
		</filter>
	</defs>
`, color)
	}
	tops["glowing-bulb-2"] = func(color string) string {
		return fmt.Sprintf(`
	<g filter="url(#topGlowingBulb02Filter0)">
		<path fill-rule="evenodd" clip-rule="evenodd" d="M30 33C30 21.9543 38.9543 13 50 13V13C61.0457 13 70 21.9543 70 33V44H30V33Z" fill="white" fill-opacity="0.3"/>
	</g>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M50 36C52.2091 36 54 35.028 54 31.7143C54 28.4006 52.2091 24 50 24C47.7909 24 46 28.4006 46 31.7143C46 35.028 47.7909 36 50 36Z" fill="white" fill-opacity="0.6"/>
	<path d="M50 14.5C49.4477 14.5 49 14.9477 49 15.5C49 16.0523 49.4477 16.5 50 16.5V14.5ZM61.6941 21.6875C62.0649 22.0968 62.6973 22.1281 63.1066 21.7573C63.5159 21.3865 63.5471 20.7541 63.1763 20.3448L61.6941 21.6875ZM65.7595 24.0473C65.5035 23.5579 64.8993 23.3686 64.4099 23.6246C63.9205 23.8806 63.7313 24.4848 63.9873 24.9742L65.7595 24.0473ZM65.4248 28.9559C65.5404 29.4959 66.0719 29.84 66.6119 29.7244C67.152 29.6088 67.4961 29.0773 67.3805 28.5373L65.4248 28.9559ZM50 16.5C54.6375 16.5 58.8065 18.4999 61.6941 21.6875L63.1763 20.3448C59.9256 16.7563 55.2256 14.5 50 14.5V16.5ZM63.9873 24.9742C64.6357 26.2139 65.1239 27.5501 65.4248 28.9559L67.3805 28.5373C67.0411 26.9518 66.4904 25.4448 65.7595 24.0473L63.9873 24.9742Z" fill="white"/>
	<rect x="20" y="36" width="60" height="16" rx="1" fill="%s"/>
	<defs>
		<filter id="topGlowingBulb02Filter0" x="22" y="5" width="56" height="47" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
			<feFlood flood-opacity="0" result="BackgroundImageFix"/>
			<feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/>
			<feOffset/>
			<feGaussianBlur stdDeviation="4"/>
			<feColorMatrix type="matrix" values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 0.5 0"/>
			<feBlend mode="normal" in2="BackgroundImageFix" result="effect1_dropShadow"/>
			<feBlend mode="normal" in="SourceGraphic" in2="effect1_dropShadow" result="shape"/>
			<feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" result="hardAlpha"/>
			<feOffset/>
			<feGaussianBlur stdDeviation="2"/>
			<feComposite in2="hardAlpha" operator="arithmetic" k2="-1" k3="1"/>
			<feColorMatrix type="matrix" values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 0.5 0"/>
			<feBlend mode="normal" in2="shape" result="effect2_innerShadow"/>
		</filter>
	</defs>
`, color)
	}
	tops["horns"] = func(color string) string {
		return fmt.Sprintf(`
	<path fill-rule="evenodd" clip-rule="evenodd" d="M71.2104 40C78.8499 33.2931 84.6313 20.6882 84 14C83.8635 12.5535 85.9998 12.2993 87 14C91.418 21.5124 89.7172 36.0672 89.1535 40H92V52H66V40H71.2104ZM16.521 13.7408C16.521 21.2733 21.4918 33.445 29.2618 40H34V52H8V40H11.2251C10.6299 36.4414 8.52929 21.6012 13.4337 14.1009C14.3353 12.7219 16.521 12.6807 16.521 13.7408Z" fill="#E1E6E8"/>
	<mask id="topHornsMask0" mask-type="alpha" maskUnits="userSpaceOnUse" x="8" y="12" width="84" height="40">
		<path fill-rule="evenodd" clip-rule="evenodd" d="M71.2104 40C78.8499 33.2931 84.6313 20.6882 84 14C83.8635 12.5535 85.9998 12.2993 87 14C91.418 21.5124 89.7172 36.0672 89.1535 40H92V52H66V40H71.2104ZM16.521 13.7408C16.521 21.2733 21.4918 33.445 29.2618 40H34V52H8V40H11.2251C10.6299 36.4414 8.52929 21.6012 13.4337 14.1009C14.3353 12.7219 16.521 12.6807 16.521 13.7408Z" fill="white"/>
	</mask>
	<g mask="url(#topHornsMask0)">
		<rect width="100" height="52" fill="%s"/>
		<rect y="40" width="100" height="12" fill="black" fill-opacity="0.4"/>
		<path fill-rule="evenodd" clip-rule="evenodd" d="M15.4558 13H31.5689V40H20.8201C13.3712 32.1499 15.4558 13 15.4558 13Z" fill="white" fill-opacity="0.4"/>
		<path fill-rule="evenodd" clip-rule="evenodd" d="M84.8203 13H92.5691V40H81.8203C87.5713 32.1946 84.8203 13 84.8203 13Z" fill="white" fill-opacity="0.4"/>
	</g>
`, color)
	}
	tops["lights"] = func(color string) string {
		return fmt.Sprintf(`
	<path fill-rule="evenodd" clip-rule="evenodd" d="M23 22C20.2386 22 18 24.2386 18 27V40H12C10.8954 40 10 40.8954 10 42V52H18H34H42H58H66H82H90V42C90 40.8954 89.1046 40 88 40H82V27C82 24.2386 79.7614 22 77 22H71C68.2386 22 66 24.2386 66 27V40H58V27C58 24.2386 55.7614 22 53 22H47C44.2386 22 42 24.2386 42 27V40H34V27C34 24.2386 31.7614 22 29 22H23Z" fill="#E1E6E8"/>
	<mask id="topLightsMask0" mask-type="alpha" maskUnits="userSpaceOnUse" x="10" y="22" width="80" height="30">
		<path fill-rule="evenodd" clip-rule="evenodd" d="M23 22C20.2386 22 18 24.2386 18 27V40H12C10.8954 40 10 40.8954 10 42V52H18H34H42H58H66H82H90V42C90 40.8954 89.1046 40 88 40H82V27C82 24.2386 79.7614 22 77 22H71C68.2386 22 66 24.2386 66 27V40H58V27C58 24.2386 55.7614 22 53 22H47C44.2386 22 42 24.2386 42 27V40H34V27C34 24.2386 31.7614 22 29 22H23Z" fill="white"/>
	</mask>
	<g mask="url(#topLightsMask0)">
		<rect width="100" height="52" fill="%s"/>
		<rect width="100" height="40" fill="white" fill-opacity="0.6"/>
		<rect x="24" y="28" width="4" height="8" rx="2" fill="white" fill-opacity="0.6"/>
		<rect x="48" y="28" width="4" height="8" rx="2" fill="white" fill-opacity="0.6"/>
		<rect x="72" y="28" width="4" height="8" rx="2" fill="white" fill-opacity="0.6"/>
	</g>
`, color)
	}
	tops["pyramid"] = func(color string) string {
		return fmt.Sprintf(`
	<path fill-rule="evenodd" clip-rule="evenodd" d="M50 8L82 52H18L50 8Z" fill="#E1E6E8"/>
	<mask id="topPyramidMask0" mask-type="alpha" maskUnits="userSpaceOnUse" x="18" y="8" width="64" height="44">
		<path fill-rule="evenodd" clip-rule="evenodd" d="M50 8L82 52H18L50 8Z" fill="white"/>
	</mask>
	<g mask="url(#topPyramidMask0)">
		<rect width="100" height="52" fill="%s"/>
		<rect x="50" y="4" width="30" height="48" fill="white" fill-opacity="0.8"/>
	</g>
`, color)
	}
	tops["radar"] = func(color string) string {
		return fmt.Sprintf(`
	<path fill-rule="evenodd" clip-rule="evenodd" d="M43.7993 28.3969C35.9888 20.5865 35.9888 7.92316 43.7993 0.112671L57.2343 13.5477L63.6874 7.09463C62.7814 5.56072 62.9874 3.55192 64.3054 2.23399C65.8675 0.671894 68.4001 0.671894 69.9622 2.23399C71.5243 3.79609 71.5243 6.32875 69.9622 7.89085C68.6443 9.20878 66.6355 9.41478 65.1016 8.50884L58.6485 14.9619L72.0835 28.3969C66.6332 33.8472 58.8199 35.4942 51.9414 33.3379V52.1127H47.9414V31.581C46.4606 30.7252 45.0661 29.6638 43.7993 28.3969Z" fill="#E1E6E8"/>
	<mask id="topRadarMask0" mask-type="alpha" maskUnits="userSpaceOnUse" x="37" y="0" width="36" height="53">
		<path fill-rule="evenodd" clip-rule="evenodd" d="M43.7993 28.3969C35.9888 20.5865 35.9888 7.92316 43.7993 0.112671L57.2343 13.5477L63.6874 7.09463C62.7814 5.56072 62.9874 3.55192 64.3054 2.23399C65.8675 0.671894 68.4001 0.671894 69.9622 2.23399C71.5243 3.79609 71.5243 6.32875 69.9622 7.89085C68.6443 9.20878 66.6355 9.41478 65.1016 8.50884L58.6485 14.9619L72.0835 28.3969C66.6332 33.8472 58.8199 35.4942 51.9414 33.3379V52.1127H47.9414V31.581C46.4606 30.7252 45.0661 29.6638 43.7993 28.3969Z" fill="white"/>
	</mask>
	<g mask="url(#topRadarMask0)">
		<rect width="100" height="52" fill="%s"/>
		<path fill-rule="evenodd" clip-rule="evenodd" d="M43.7988 0.112671C35.9883 7.92316 35.9883 20.5865 43.7988 28.3969C51.6093 36.2074 64.2726 36.2074 72.0831 28.3969" fill="white" fill-opacity="0.2"/>
		<path fill-rule="evenodd" clip-rule="evenodd" d="M64.3054 7.89092C65.8675 9.45302 68.4001 9.45302 69.9622 7.89092C71.5243 6.32882 71.5243 3.79616 69.9622 2.23407C68.4001 0.67197 65.8675 0.67197 64.3054 2.23407C62.7433 3.79616 62.7433 6.32882 64.3054 7.89092Z" fill="white" fill-opacity="0.8"/>
	</g>
`, color)
	}

	topNames = make([]string, 0, len(tops))
	for k := range tops {
		topNames = append(topNames, k)
	}
	sort.Strings(topNames)
}
