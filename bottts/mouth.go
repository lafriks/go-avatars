package bottts

import (
	"sort"
)

var mouth = map[string]string{
	"bite": `
	<rect x="4" y="5" width="68" height="22" rx="5" fill="black" fill-opacity="0.2"/>
	<rect x="8" y="9" width="60" height="14" rx="2" fill="black" fill-opacity="0.6"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M20 17L26 9H14L20 17Z" fill="#E1E6E8"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M32 17L38 9H26L32 17Z" fill="#E1E6E8"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M44 17L50 9H38L44 17Z" fill="#E1E6E8"/>
	<path fill-rule="evenodd" clip-rule="evenodd" d="M56 17L62 9H50L56 17Z" fill="#E1E6E8"/>
`,
	"diagram": `
	<rect x="4" y="4" width="68" height="24" rx="5" fill="black" fill-opacity="0.2"/>
	<rect x="8" y="8" width="60" height="16" rx="2" fill="black" fill-opacity="0.8"/>
	<path d="M9 17H20L22 13L25 20L29 12L31 21L34 10L37 20L40 17H55L58 13L60 20L63 17H67" stroke="#4EFAC9" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
`,
	"grill-1": `
	<rect x="12" y="12" width="4" height="8" rx="2" fill="black" fill-opacity="0.6"/>
	<rect x="36" y="12" width="4" height="8" rx="2" fill="black" fill-opacity="0.6"/>
	<rect x="24" y="12" width="4" height="8" rx="2" fill="black" fill-opacity="0.6"/>
	<rect x="48" y="12" width="4" height="8" rx="2" fill="black" fill-opacity="0.6"/>
	<rect x="60" y="12" width="4" height="8" rx="2" fill="black" fill-opacity="0.6"/>
`,
	"grill-2": `
	<rect x="28" y="10" width="6" height="14" rx="2" fill="black" fill-opacity="0.6"/>
	<rect x="14" y="10" width="6" height="14" rx="2" fill="black" fill-opacity="0.6"/>
	<rect x="42" y="10" width="6" height="14" rx="2" fill="black" fill-opacity="0.6"/>
	<rect x="56" y="10" width="6" height="14" rx="2" fill="black" fill-opacity="0.6"/>
`,
	"grill-3": `
	<rect x="4" y="5" width="68" height="22" rx="5" fill="black" fill-opacity="0.2"/>
	<rect x="8" y="9" width="60" height="14" rx="2" fill="white"/>
	<rect x="18" y="9" width="4" height="14" fill="black" fill-opacity="0.1"/>
	<rect x="42" y="9" width="4" height="14" fill="black" fill-opacity="0.1"/>
	<rect x="30" y="9" width="4" height="14" fill="black" fill-opacity="0.1"/>
	<rect x="54" y="9" width="4" height="14" fill="black" fill-opacity="0.1"/>
`,
	"smile-1": `
	<path d="M27.0493 8.44151C26.8055 7.36419 27.4811 6.29318 28.5584 6.04935C29.6358 5.80551 30.7068 6.48119 30.9506 7.55851C31.72 10.9578 34.4016 13 37.9999 13C41.5983 13 44.2799 10.9578 45.0493 7.55851C45.2931 6.48119 46.3641 5.80551 47.4414 6.04935C48.5188 6.29318 49.1944 7.36419 48.9506 8.44151C47.7599 13.7024 43.4298 17 37.9999 17C32.5701 17 28.24 13.7024 27.0493 8.44151Z" fill="black" fill-opacity="0.6"/>
`,
	"smile-2": `
	<path fill-rule="evenodd" clip-rule="evenodd" d="M18 10.2222C18 21.7845 24.4741 28 38 28C51.5182 28 58 21.6615 58 10.2222C58 9.49622 57.1739 8 55 8C39.2707 8 29.1917 8 21 8C18.949 8 18 9.38479 18 10.2222Z" fill="black" fill-opacity="0.8"/>
	<mask id="mouthSmilie02Mask0" mask-type="alpha" maskUnits="userSpaceOnUse" x="18" y="8" width="40" height="20">
		<path fill-rule="evenodd" clip-rule="evenodd" d="M18 10.2222C18 21.7845 24.4741 28 38 28C51.5182 28 58 21.6615 58 10.2222C58 9.49622 57.1739 8 55 8C39.2707 8 29.1917 8 21 8C18.949 8 18 9.38479 18 10.2222Z" fill="white"/>
	</mask>
	<g mask="url(#mouthSmilie02Mask0)">
		<rect x="30" y="2" width="16" height="14" rx="2" fill="white"/>
	</g>
`,
	"square-1": `
	<rect x="24" y="6" width="27" height="8" rx="4" fill="black" fill-opacity="0.8"/>
`,
	"square-2": `
	<rect x="16" y="8" width="44" height="4" rx="2" fill="black" fill-opacity="0.8"/>
`,
}

var mouthNames []string

func init() {
	mouthNames = make([]string, 0, len(mouth))
	for k := range mouth {
		mouthNames = append(mouthNames, k)
	}
	sort.Strings(mouthNames)
}
