package main

import "fmt"

func main() {
	var nums = []int{387, 74, 244, 281, 942, 3, 735, 363, 170, 89, 541, 563, 504, 698, 609, 853, 374, 859, 166, 89, 89, 877, 51, 802, 758, 829, 528, 21, 242, 227, 438, 940, 622, 601, 40, 552, 839, 179, 473, 206, 196, 903, 505, 721, 919, 526, 570, 47, 888, 408, 453, 738, 412, 570, 998, 316, 569, 928, 549, 155, 995, 99, 52, 221, 262, 622, 811, 100, 477, 795, 868, 369, 331, 16, 507, 719, 311, 123, 575, 275, 548, 468, 703, 488, 736, 885, 13, 985, 742, 646, 651, 594, 181, 74, 87, 217, 266, 395, 773, 710, 771, 923, 492, 821, 236, 836, 468, 775, 772, 34, 93, 904, 504, 695, 756, 607, 698, 65, 957, 824, 55, 954, 246, 733, 84, 486, 290, 510, 6, 250, 95, 486, 810, 79, 197, 571, 198, 310, 522, 503, 49, 186, 65, 655, 532, 230, 868, 324, 135, 334, 78, 305, 158, 973, 629, 294, 611, 95, 313, 815, 161, 865, 532, 33, 917, 298, 863, 128, 476, 726, 458, 408, 608, 359, 42, 219, 926, 654, 851, 504, 846, 501, 429, 404, 221, 387, 462, 767, 336, 301, 74, 453, 222, 65, 929, 933, 888, 751, 25, 173, 787, 547, 823, 53, 68, 543, 684, 343, 780, 903, 324, 958, 33, 939, 562, 561, 895, 432, 826, 936, 131, 227, 469, 275, 914, 951, 489, 788, 637, 833, 84, 700, 646, 509, 57, 653, 268, 85, 736, 476, 198, 977, 920, 792, 166, 216, 855, 793, 378, 703, 20, 737, 673, 93, 233, 915, 694, 181, 446, 779, 645, 239, 305, 588, 781, 482, 700, 74, 294, 718, 352, 247, 465, 227, 788, 328, 414, 948, 859, 354, 407, 262, 216, 401, 60, 389, 844, 590, 556, 935, 492, 324, 925, 865, 531, 603, 871, 266, 731, 13, 84, 177, 684, 550, 615, 348, 975, 733, 155, 453, 113, 554, 184, 881, 647, 260, 551, 158, 938, 245, 182, 653, 861, 539, 520, 103, 570, 461, 786, 128, 789, 462, 751, 162, 84, 110, 703, 829, 99, 575, 295, 886, 18, 767, 440, 473, 949, 121, 598, 112, 447, 595, 79, 303, 923, 275, 287, 658, 794, 811, 641, 812, 209, 182, 979, 76, 939, 525, 492, 848, 170, 968, 516, 458, 513, 479, 918, 453, 563, 344, 260, 493, 964, 788, 983, 601, 580, 793, 549, 197, 687, 198, 200, 245, 905, 549, 91, 875, 199, 565, 177, 3, 184, 493, 978, 514, 432, 787, 903, 47, 334, 505, 823, 744, 871, 981, 510, 981, 21, 802, 918, 326, 715, 605, 607, 395, 571, 646, 748, 324, 895, 586, 50, 868, 613, 59, 604, 660, 799, 488, 84, 129, 788, 446, 43, 551, 398, 810, 929, 368, 686, 751, 997, 632, 899, 376, 591, 281, 492, 470, 330, 716, 923, 124, 591, 453, 458, 674, 834, 145, 682, 456, 639, 689, 894, 236, 667, 378, 255, 406, 678, 866, 777, 67, 546, 170, 898, 456, 471, 998, 732, 683, 713, 117, 33, 13, 810, 120, 125, 556, 609, 937, 579, 543, 305, 201, 637, 340, 885, 469, 298, 742, 668, 230, 895, 95, 872, 598, 802, 416, 18, 212, 376, 483, 686, 801, 124, 690, 227, 460, 84, 207, 275, 874, 627, 805, 324, 569, 832, 704, 190, 476, 520, 302, 593, 439, 205, 859, 577, 590, 406, 314, 184, 147, 275, 166, 173, 674, 722, 762, 599, 147, 916, 20, 640, 891, 74, 889, 879, 493, 521, 344, 539, 316, 109, 148, 378, 646, 217, 995, 802, 275, 141, 560, 880, 866, 603, 222, 91, 316, 635, 811, 125, 438, 835, 604, 193, 909, 997, 256, 903, 897, 950, 971, 945, 957, 406, 245, 164, 985, 705, 188, 871, 205, 483, 574, 49, 583, 119, 682, 821, 128, 565, 981, 376, 262, 763, 487, 505, 766, 501, 691, 25, 850, 50, 817, 190, 918, 241, 531, 161, 737, 490, 791, 562, 998, 247, 895, 532, 600, 816, 530, 693, 16, 49, 933, 446, 310, 522, 761, 790, 295, 537, 317, 477, 822, 102, 33, 429, 287, 601, 161, 669, 227, 60, 74, 917, 990, 387, 326, 732, 711, 742, 117, 328, 617, 408, 13, 406, 443, 412, 429, 852, 211, 567, 629, 712, 343, 363, 616, 681, 145, 145, 326, 720, 734, 995, 720, 904, 706, 348, 562, 205, 76, 805, 995, 760, 713, 980, 787, 798, 393, 65, 359, 298, 878, 104, 446, 422, 316, 435, 685, 611, 798, 708, 200, 260, 67, 232, 958, 594, 267, 916, 514, 731, 181, 538, 973, 702, 435, 525, 459, 314, 227, 182, 845, 458, 255, 378, 539, 352, 667, 86, 487, 597, 406, 39, 978, 21, 725, 393, 931, 508, 227, 779, 658, 889, 33, 499, 6, 25, 503, 49, 669, 348, 131, 227, 161, 966, 763, 780, 278, 145, 599, 597, 960, 198, 496, 275, 186, 438, 143, 21, 629, 877, 33, 137, 318, 337, 785, 703, 128, 387, 833, 201, 135, 885, 268, 340, 809, 287, 246, 187, 28, 374, 805, 851, 493, 266, 28, 830, 726, 841, 634, 260, 694, 507, 575, 50, 543, 940, 3, 738, 763, 915, 961, 285, 947, 74, 809, 148, 781, 841, 456, 880, 569, 236, 292, 996, 356, 644, 636, 736, 247, 59, 216, 611, 129, 247, 109, 4, 416, 539, 572, 344, 252, 255, 737, 407, 147, 60, 877, 448, 841, 891, 843, 28, 236, 196, 326, 448, 363, 227, 128, 406, 657, 307, 451, 307, 99, 519, 937, 412, 799, 705, 131, 404, 33, 952, 164, 3, 985, 187, 190, 811, 819, 962, 390, 488, 340, 227, 275, 343, 817, 3, 634, 779, 109, 397, 149, 399, 55, 747, 173, 250, 869, 896, 568, 916, 578, 84, 865, 42, 547, 549, 158, 966, 497, 969, 369, 166, 352, 340, 209, 479, 701, 355, 382, 809, 665, 278, 25, 665, 524, 135, 746, 563, 15, 164, 592, 82, 600, 691, 537, 369, 704, 170, 46, 336, 841, 801, 212, 657, 834, 496, 109, 893, 216, 571, 59, 211, 67, 465, 729, 619, 691, 316, 793, 52, 426, 13, 407, 797, 885, 940, 807, 996, 103, 977, 472, 804, 746, 884, 174, 972, 161, 95, 447, 438, 198, 702, 772, 78, 721, 577, 491, 583, 782, 354, 960, 759, 820, 487, 34, 997, 161, 533, 117, 522, 758, 87, 998, 716, 736, 28, 141, 401, 992, 252, 369, 728, 25, 137, 404, 889, 634, 57, 784, 441, 416, 113, 523, 97, 120, 451, 42, 469, 473, 538, 387, 922, 940, 913, 962, 577, 997, 390, 275, 573, 227, 179, 987, 701, 174, 919, 479, 18, 316, 693, 677, 786, 206, 618, 275, 552, 663, 451, 446, 6, 59, 889, 117, 20, 698, 878, 651, 281, 59, 605, 541, 500, 256, 670, 260, 986, 673, 235, 257, 645, 481, 250, 456, 399, 438, 759, 473, 530, 882, 498, 667, 164, 117, 468, 458, 78, 880, 811, 703, 841, 568, 18, 656, 21, 911, 886, 718, 25, 221, 937, 362, 103, 250, 275, 188, 352, 98, 28, 221, 435, 324, 330, 119, 628, 278, 913, 389, 28, 971, 294, 397, 3}
	var target = 470
	TwoSum(nums, target)
	var a = TwoSum(nums, target)
	if a != nil {
		for _, value := range a {
			fmt.Println(value)
			fmt.Println(nums[value])
		}
	}

}

func TwoSum(nums []int, target int) []int {
	var lennum = len(nums) - 1
	for i := 0; i <= lennum; i++ {
		for j := 1; j <= lennum; j++ {
			if i == j {
				j = j + 1

			}
			if j > lennum {
				j = lennum
			}

			var sum = nums[i] + nums[j]

			if sum == target {

				var reust = []int{i, j}
				return reust
			}
		}
	}
	return nil
}
