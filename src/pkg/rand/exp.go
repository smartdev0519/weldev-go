// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

import (
	"math";
)

/*
 * Exponential distribution
 *
 * See "The Ziggurat Method for Generating Random Variables"
 * (Marsaglia & Tsang, 2000)
 * http://www.jstatsoft.org/v05/i08/paper [pdf]
 */

const (
	re = 7.69711747013104972;
)

// ExpFloat64 returns an exponentially distributed float64 in the range
// (0, +math.MaxFloat64] with an exponential distribution whose rate parameter
// (lambda) is 1 and whose mean is 1/lambda (1).
// To produce a distribution with a different rate parameter,
// callers can adjust the output using:
//
//  sample = ExpFloat64() / desiredRateParameter
//
func (r *Rand) ExpFloat64() float64 {
	for {
		j := r.Uint32();
		i := j&0xFF;
		x := float64(j)*float64(we[i]);
		if j < ke[i] {
			return x;
		}
		if i == 0 {
			return re - math.Log(r.Float64());
		}
		if fe[i] + float32(r.Float64())*(fe[i-1]-fe[i]) < float32(math.Exp(-x)) {
			return x;
		}
	}
	panic("unreachable");
}

var ke = [256]uint32{
	0xe290a139, 0x0, 0x9beadebc, 0xc377ac71, 0xd4ddb990,
	0xde893fb8, 0xe4a8e87c, 0xe8dff16a, 0xebf2deab, 0xee49a6e8,
	0xf0204efd, 0xf19bdb8e, 0xf2d458bb, 0xf3da104b, 0xf4b86d78,
	0xf577ad8a, 0xf61de83d, 0xf6afb784, 0xf730a573, 0xf7a37651,
	0xf80a5bb6, 0xf867189d, 0xf8bb1b4f, 0xf9079062, 0xf94d70ca,
	0xf98d8c7d, 0xf9c8928a, 0xf9ff175b, 0xfa319996, 0xfa6085f8,
	0xfa8c3a62, 0xfab5084e, 0xfadb36c8, 0xfaff0410, 0xfb20a6ea,
	0xfb404fb4, 0xfb5e2951, 0xfb7a59e9, 0xfb95038c, 0xfbae44ba,
	0xfbc638d8, 0xfbdcf892, 0xfbf29a30, 0xfc0731df, 0xfc1ad1ed,
	0xfc2d8b02, 0xfc3f6c4d, 0xfc5083ac, 0xfc60ddd1, 0xfc708662,
	0xfc7f8810, 0xfc8decb4, 0xfc9bbd62, 0xfca9027c, 0xfcb5c3c3,
	0xfcc20864, 0xfccdd70a, 0xfcd935e3, 0xfce42ab0, 0xfceebace,
	0xfcf8eb3b, 0xfd02c0a0, 0xfd0c3f59, 0xfd156b7b, 0xfd1e48d6,
	0xfd26daff, 0xfd2f2552, 0xfd372af7, 0xfd3eeee5, 0xfd4673e7,
	0xfd4dbc9e, 0xfd54cb85, 0xfd5ba2f2, 0xfd62451b, 0xfd68b415,
	0xfd6ef1da, 0xfd750047, 0xfd7ae120, 0xfd809612, 0xfd8620b4,
	0xfd8b8285, 0xfd90bcf5, 0xfd95d15e, 0xfd9ac10b, 0xfd9f8d36,
	0xfda43708, 0xfda8bf9e, 0xfdad2806, 0xfdb17141, 0xfdb59c46,
	0xfdb9a9fd, 0xfdbd9b46, 0xfdc170f6, 0xfdc52bd8, 0xfdc8ccac,
	0xfdcc542d, 0xfdcfc30b, 0xfdd319ef, 0xfdd6597a, 0xfdd98245,
	0xfddc94e5, 0xfddf91e6, 0xfde279ce, 0xfde54d1f, 0xfde80c52,
	0xfdeab7de, 0xfded5034, 0xfdefd5be, 0xfdf248e3, 0xfdf4aa06,
	0xfdf6f984, 0xfdf937b6, 0xfdfb64f4, 0xfdfd818d, 0xfdff8dd0,
	0xfe018a08, 0xfe03767a, 0xfe05536c, 0xfe07211c, 0xfe08dfc9,
	0xfe0a8fab, 0xfe0c30fb, 0xfe0dc3ec, 0xfe0f48b1, 0xfe10bf76,
	0xfe122869, 0xfe1383b4, 0xfe14d17c, 0xfe1611e7, 0xfe174516,
	0xfe186b2a, 0xfe19843e, 0xfe1a9070, 0xfe1b8fd6, 0xfe1c8289,
	0xfe1d689b, 0xfe1e4220, 0xfe1f0f26, 0xfe1fcfbc, 0xfe2083ed,
	0xfe212bc3, 0xfe21c745, 0xfe225678, 0xfe22d95f, 0xfe234ffb,
	0xfe23ba4a, 0xfe241849, 0xfe2469f2, 0xfe24af3c, 0xfe24e81e,
	0xfe25148b, 0xfe253474, 0xfe2547c7, 0xfe254e70, 0xfe25485a,
	0xfe25356a, 0xfe251586, 0xfe24e88f, 0xfe24ae64, 0xfe2466e1,
	0xfe2411df, 0xfe23af34, 0xfe233eb4, 0xfe22c02c, 0xfe22336b,
	0xfe219838, 0xfe20ee58, 0xfe20358c, 0xfe1f6d92, 0xfe1e9621,
	0xfe1daef0, 0xfe1cb7ac, 0xfe1bb002, 0xfe1a9798, 0xfe196e0d,
	0xfe1832fd, 0xfe16e5fe, 0xfe15869d, 0xfe141464, 0xfe128ed3,
	0xfe10f565, 0xfe0f478c, 0xfe0d84b1, 0xfe0bac36, 0xfe09bd73,
	0xfe07b7b5, 0xfe059a40, 0xfe03644c, 0xfe011504, 0xfdfeab88,
	0xfdfc26e9, 0xfdf98629, 0xfdf6c83b, 0xfdf3ec01, 0xfdf0f04a,
	0xfdedd3d1, 0xfdea953d, 0xfde7331e, 0xfde3abe9, 0xfddffdfb,
	0xfddc2791, 0xfdd826cd, 0xfdd3f9a8, 0xfdcf9dfc, 0xfdcb1176,
	0xfdc65198, 0xfdc15bb3, 0xfdbc2ce2, 0xfdb6c206, 0xfdb117be,
	0xfdab2a63, 0xfda4f5fd, 0xfd9e7640, 0xfd97a67a, 0xfd908192,
	0xfd8901f2, 0xfd812182, 0xfd78d98e, 0xfd7022bb, 0xfd66f4ed,
	0xfd5d4732, 0xfd530f9c, 0xfd48432b, 0xfd3cd59a, 0xfd30b936,
	0xfd23dea4, 0xfd16349e, 0xfd07a7a3, 0xfcf8219b, 0xfce7895b,
	0xfcd5c220, 0xfcc2aadb, 0xfcae1d5e, 0xfc97ed4e, 0xfc7fe6d4,
	0xfc65ccf3, 0xfc495762, 0xfc2a2fc8, 0xfc07ee19, 0xfbe213c1,
	0xfbb8051a, 0xfb890078, 0xfb5411a5, 0xfb180005, 0xfad33482,
	0xfa839276, 0xfa263b32, 0xf9b72d1c, 0xf930a1a2, 0xf889f023,
	0xf7b577d2, 0xf69c650c, 0xf51530f0, 0xf2cb0e3c, 0xeeefb15d,
	0xe6da6ecf,
}
var we = [256]float32{
	2.0249555e-09, 1.486674e-11, 2.4409617e-11, 3.1968806e-11,
	3.844677e-11, 4.4228204e-11, 4.9516443e-11, 5.443359e-11,
	5.905944e-11, 6.344942e-11, 6.7643814e-11, 7.1672945e-11,
	7.556032e-11, 7.932458e-11, 8.298079e-11, 8.654132e-11,
	9.0016515e-11, 9.3415074e-11, 9.674443e-11, 1.0001099e-10,
	1.03220314e-10, 1.06377254e-10, 1.09486115e-10, 1.1255068e-10,
	1.1557435e-10, 1.1856015e-10, 1.2151083e-10, 1.2442886e-10,
	1.2731648e-10, 1.3017575e-10, 1.3300853e-10, 1.3581657e-10,
	1.3860142e-10, 1.4136457e-10, 1.4410738e-10, 1.4683108e-10,
	1.4953687e-10, 1.5222583e-10, 1.54899e-10, 1.5755733e-10,
	1.6020171e-10, 1.6283301e-10, 1.6545203e-10, 1.6805951e-10,
	1.7065617e-10, 1.732427e-10, 1.7581973e-10, 1.7838787e-10,
	1.8094774e-10, 1.8349985e-10, 1.8604476e-10, 1.8858298e-10,
	1.9111498e-10, 1.9364126e-10, 1.9616223e-10, 1.9867835e-10,
	2.0119004e-10, 2.0369768e-10, 2.0620168e-10, 2.087024e-10,
	2.1120022e-10, 2.136955e-10, 2.1618855e-10, 2.1867974e-10,
	2.2116936e-10, 2.2365775e-10, 2.261452e-10, 2.2863202e-10,
	2.311185e-10, 2.3360494e-10, 2.360916e-10, 2.3857874e-10,
	2.4106667e-10, 2.4355562e-10, 2.4604588e-10, 2.485377e-10,
	2.5103128e-10, 2.5352695e-10, 2.560249e-10, 2.585254e-10,
	2.6102867e-10, 2.6353494e-10, 2.6604446e-10, 2.6855745e-10,
	2.7107416e-10, 2.7359479e-10, 2.761196e-10, 2.7864877e-10,
	2.8118255e-10, 2.8372119e-10, 2.8626485e-10, 2.888138e-10,
	2.9136826e-10, 2.939284e-10, 2.9649452e-10, 2.9906677e-10,
	3.016454e-10, 3.0423064e-10, 3.0682268e-10, 3.0942177e-10,
	3.1202813e-10, 3.1464195e-10, 3.1726352e-10, 3.19893e-10,
	3.2253064e-10, 3.251767e-10, 3.2783135e-10, 3.3049485e-10,
	3.3316744e-10, 3.3584938e-10, 3.3854083e-10, 3.4124212e-10,
	3.4395342e-10, 3.46675e-10, 3.4940711e-10, 3.5215003e-10,
	3.5490397e-10, 3.5766917e-10, 3.6044595e-10, 3.6323455e-10,
	3.660352e-10, 3.6884823e-10, 3.7167386e-10, 3.745124e-10,
	3.773641e-10, 3.802293e-10, 3.8310827e-10, 3.860013e-10,
	3.8890866e-10, 3.918307e-10, 3.9476775e-10, 3.9772008e-10,
	4.0068804e-10, 4.0367196e-10, 4.0667217e-10, 4.09689e-10,
	4.1272286e-10, 4.1577405e-10, 4.1884296e-10, 4.2192994e-10,
	4.250354e-10, 4.281597e-10, 4.313033e-10, 4.3446652e-10,
	4.3764986e-10, 4.408537e-10, 4.4407847e-10, 4.4732465e-10,
	4.5059267e-10, 4.5388301e-10, 4.571962e-10, 4.6053267e-10,
	4.6389292e-10, 4.6727755e-10, 4.70687e-10, 4.741219e-10,
	4.7758275e-10, 4.810702e-10, 4.845848e-10, 4.8812715e-10,
	4.9169796e-10, 4.9529775e-10, 4.989273e-10, 5.0258725e-10,
	5.0627835e-10, 5.100013e-10, 5.1375687e-10, 5.1754584e-10,
	5.21369e-10, 5.2522725e-10, 5.2912136e-10, 5.330522e-10,
	5.370208e-10, 5.4102806e-10, 5.45075e-10, 5.491625e-10,
	5.532918e-10, 5.5746385e-10, 5.616799e-10, 5.6594107e-10,
	5.7024857e-10, 5.746037e-10, 5.7900773e-10, 5.834621e-10,
	5.8796823e-10, 5.925276e-10, 5.971417e-10, 6.018122e-10,
	6.065408e-10, 6.113292e-10, 6.1617933e-10, 6.2109295e-10,
	6.260722e-10, 6.3111916e-10, 6.3623595e-10, 6.4142497e-10,
	6.4668854e-10, 6.5202926e-10, 6.5744976e-10, 6.6295286e-10,
	6.6854156e-10, 6.742188e-10, 6.79988e-10, 6.858526e-10,
	6.9181616e-10, 6.978826e-10, 7.04056e-10, 7.103407e-10,
	7.167412e-10, 7.2326256e-10, 7.2990985e-10, 7.366886e-10,
	7.4360473e-10, 7.5066453e-10, 7.5787476e-10, 7.6524265e-10,
	7.7277595e-10, 7.80483e-10, 7.883728e-10, 7.9645507e-10,
	8.047402e-10, 8.1323964e-10, 8.219657e-10, 8.309319e-10,
	8.401528e-10, 8.496445e-10, 8.594247e-10, 8.6951274e-10,
	8.799301e-10, 8.9070046e-10, 9.018503e-10, 9.134092e-10,
	9.254101e-10, 9.378904e-10, 9.508923e-10, 9.644638e-10,
	9.786603e-10, 9.935448e-10, 1.0091913e-09, 1.025686e-09,
	1.0431306e-09, 1.0616465e-09, 1.08138e-09, 1.1025096e-09,
	1.1252564e-09, 1.1498986e-09, 1.1767932e-09, 1.206409e-09,
	1.2393786e-09, 1.276585e-09, 1.3193139e-09, 1.3695435e-09,
	1.4305498e-09, 1.508365e-09, 1.6160854e-09, 1.7921248e-09,
}
var fe = [256]float32{
	1, 0.9381437, 0.90046996, 0.87170434, 0.8477855, 0.8269933,
	0.8084217, 0.7915276, 0.77595687, 0.7614634, 0.7478686,
	0.7350381, 0.72286767, 0.71127474, 0.70019263, 0.6895665,
	0.67935055, 0.6695063, 0.66000086, 0.65080583, 0.6418967,
	0.63325197, 0.6248527, 0.6166822, 0.60872537, 0.60096896,
	0.5934009, 0.58601034, 0.5787874, 0.57172304, 0.5648092,
	0.5580383, 0.5514034, 0.5448982, 0.5385169, 0.53225386,
	0.5261042, 0.52006316, 0.5141264, 0.50828975, 0.5025495,
	0.496902, 0.49134386, 0.485872, 0.48048335, 0.4751752,
	0.46994483, 0.46478975, 0.45970762, 0.45469615, 0.44975325,
	0.44487688, 0.44006512, 0.43531612, 0.43062815, 0.42599955,
	0.42142874, 0.4169142, 0.41245446, 0.40804818, 0.403694,
	0.3993907, 0.39513698, 0.39093173, 0.38677382, 0.38266218,
	0.37859577, 0.37457356, 0.37059465, 0.3666581, 0.362763,
	0.35890847, 0.35509375, 0.351318, 0.3475805, 0.34388044,
	0.34021714, 0.3365899, 0.33299807, 0.32944095, 0.32591796,
	0.3224285, 0.3189719, 0.31554767, 0.31215525, 0.30879408,
	0.3054636, 0.3021634, 0.29889292, 0.2956517, 0.29243928,
	0.28925523, 0.28609908, 0.28297043, 0.27986884, 0.27679393,
	0.2737453, 0.2707226, 0.2677254, 0.26475343, 0.26180625,
	0.25888354, 0.25598502, 0.2531103, 0.25025907, 0.24743107,
	0.24462597, 0.24184346, 0.23908329, 0.23634516, 0.23362878,
	0.23093392, 0.2282603, 0.22560766, 0.22297576, 0.22036438,
	0.21777324, 0.21520215, 0.21265087, 0.21011916, 0.20760682,
	0.20511365, 0.20263945, 0.20018397, 0.19774707, 0.19532852,
	0.19292815, 0.19054577, 0.1881812, 0.18583426, 0.18350479,
	0.1811926, 0.17889754, 0.17661946, 0.17435817, 0.17211354,
	0.1698854, 0.16767362, 0.16547804, 0.16329853, 0.16113494,
	0.15898713, 0.15685499, 0.15473837, 0.15263714, 0.15055119,
	0.14848037, 0.14642459, 0.14438373, 0.14235765, 0.14034624,
	0.13834943, 0.13636707, 0.13439907, 0.13244532, 0.13050574,
	0.1285802, 0.12666863, 0.12477092, 0.12288698, 0.12101672,
	0.119160056, 0.1173169, 0.115487166, 0.11367077, 0.11186763,
	0.11007768, 0.10830083, 0.10653701, 0.10478614, 0.10304816,
	0.101323, 0.09961058, 0.09791085, 0.09622374, 0.09454919,
	0.09288713, 0.091237515, 0.08960028, 0.087975375, 0.08636274,
	0.08476233, 0.083174095, 0.081597984, 0.08003395, 0.07848195,
	0.076941945, 0.07541389, 0.07389775, 0.072393484, 0.07090106,
	0.069420435, 0.06795159, 0.066494495, 0.06504912, 0.063615434,
	0.062193416, 0.060783047, 0.059384305, 0.057997175,
	0.05662164, 0.05525769, 0.053905312, 0.052564494, 0.051235236,
	0.049917534, 0.048611384, 0.047316793, 0.046033762, 0.0447623,
	0.043502413, 0.042254124, 0.041017443, 0.039792392,
	0.038578995, 0.037377283, 0.036187284, 0.035009038,
	0.033842582, 0.032687962, 0.031545233, 0.030414443, 0.02929566,
	0.02818895, 0.027094385, 0.026012046, 0.024942026, 0.023884421,
	0.022839336, 0.021806888, 0.020787204, 0.019780423, 0.0187867,
	0.0178062, 0.016839107, 0.015885621, 0.014945968, 0.014020392,
	0.013109165, 0.012212592, 0.011331013, 0.01046481, 0.009614414,
	0.008780315, 0.007963077, 0.0071633533, 0.006381906,
	0.0056196423, 0.0048776558, 0.004157295, 0.0034602648,
	0.0027887989, 0.0021459677, 0.0015362998, 0.0009672693,
	0.00045413437,
}
