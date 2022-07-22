// WARNING: This file has automatically been generated on Fri, 22 Jul 2022 22:52:51 EEST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package all

/*
#include "DEV_Config.h"
#include "bcm2835.h"
#include "EPD_IT8951.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// EPD_RST_PIN as defined in all/DEV_Config.h:67
	EPD_RST_PIN = 17
	// EPD_CS_PIN as defined in all/DEV_Config.h:68
	EPD_CS_PIN = 8
	// EPD_BUSY_PIN as defined in all/DEV_Config.h:69
	EPD_BUSY_PIN = 24
	// BCM2835_VERSION as defined in all/bcm2835.h:635
	BCM2835_VERSION = 10071
	// BCM2835_CORE_CLK_HZ as defined in all/bcm2835.h:667
	BCM2835_CORE_CLK_HZ = 250000000
	// BCM2835_PERI_BASE as defined in all/bcm2835.h:679
	BCM2835_PERI_BASE = 536870912
	// BCM2835_PERI_SIZE as defined in all/bcm2835.h:681
	BCM2835_PERI_SIZE = 16777216
	// BCM2835_RPI2_PERI_BASE as defined in all/bcm2835.h:683
	BCM2835_RPI2_PERI_BASE = 1056964608
	// BCM2835_RPI4_PERI_BASE as defined in all/bcm2835.h:685
	BCM2835_RPI4_PERI_BASE = 4261412864
	// BCM2835_RPI4_PERI_SIZE as defined in all/bcm2835.h:687
	BCM2835_RPI4_PERI_SIZE = 25165824
	// BCM2835_ST_BASE as defined in all/bcm2835.h:692
	BCM2835_ST_BASE = 12288
	// BCM2835_GPIO_PADS as defined in all/bcm2835.h:694
	BCM2835_GPIO_PADS = 1048576
	// BCM2835_CLOCK_BASE as defined in all/bcm2835.h:696
	BCM2835_CLOCK_BASE = 1052672
	// BCM2835_GPIO_BASE as defined in all/bcm2835.h:698
	BCM2835_GPIO_BASE = 2097152
	// BCM2835_SPI0_BASE as defined in all/bcm2835.h:700
	BCM2835_SPI0_BASE = 2113536
	// BCM2835_BSC0_BASE as defined in all/bcm2835.h:702
	BCM2835_BSC0_BASE = 2117632
	// BCM2835_GPIO_PWM as defined in all/bcm2835.h:704
	BCM2835_GPIO_PWM = 2146304
	// BCM2835_AUX_BASE as defined in all/bcm2835.h:706
	BCM2835_AUX_BASE = 2183168
	// BCM2835_SPI1_BASE as defined in all/bcm2835.h:708
	BCM2835_SPI1_BASE = 2183296
	// BCM2835_SPI2_BASE as defined in all/bcm2835.h:710
	BCM2835_SPI2_BASE = 2183360
	// BCM2835_BSC1_BASE as defined in all/bcm2835.h:712
	BCM2835_BSC1_BASE = 8404992
	// BCM2835_SMI_BASE as defined in all/bcm2835.h:716
	BCM2835_SMI_BASE = 6291456
	// BCM2835_PAGE_SIZE as defined in all/bcm2835.h:807
	BCM2835_PAGE_SIZE = 4096
	// BCM2835_BLOCK_SIZE as defined in all/bcm2835.h:809
	BCM2835_BLOCK_SIZE = 4096
	// BCM2835_GPFSEL0 as defined in all/bcm2835.h:819
	BCM2835_GPFSEL0 = 0
	// BCM2835_GPFSEL1 as defined in all/bcm2835.h:820
	BCM2835_GPFSEL1 = 4
	// BCM2835_GPFSEL2 as defined in all/bcm2835.h:821
	BCM2835_GPFSEL2 = 8
	// BCM2835_GPFSEL3 as defined in all/bcm2835.h:822
	BCM2835_GPFSEL3 = 12
	// BCM2835_GPFSEL4 as defined in all/bcm2835.h:823
	BCM2835_GPFSEL4 = 16
	// BCM2835_GPFSEL5 as defined in all/bcm2835.h:824
	BCM2835_GPFSEL5 = 20
	// BCM2835_GPSET0 as defined in all/bcm2835.h:825
	BCM2835_GPSET0 = 28
	// BCM2835_GPSET1 as defined in all/bcm2835.h:826
	BCM2835_GPSET1 = 32
	// BCM2835_GPCLR0 as defined in all/bcm2835.h:827
	BCM2835_GPCLR0 = 40
	// BCM2835_GPCLR1 as defined in all/bcm2835.h:828
	BCM2835_GPCLR1 = 44
	// BCM2835_GPLEV0 as defined in all/bcm2835.h:829
	BCM2835_GPLEV0 = 52
	// BCM2835_GPLEV1 as defined in all/bcm2835.h:830
	BCM2835_GPLEV1 = 56
	// BCM2835_GPEDS0 as defined in all/bcm2835.h:831
	BCM2835_GPEDS0 = 64
	// BCM2835_GPEDS1 as defined in all/bcm2835.h:832
	BCM2835_GPEDS1 = 68
	// BCM2835_GPREN0 as defined in all/bcm2835.h:833
	BCM2835_GPREN0 = 76
	// BCM2835_GPREN1 as defined in all/bcm2835.h:834
	BCM2835_GPREN1 = 80
	// BCM2835_GPFEN0 as defined in all/bcm2835.h:835
	BCM2835_GPFEN0 = 88
	// BCM2835_GPFEN1 as defined in all/bcm2835.h:836
	BCM2835_GPFEN1 = 92
	// BCM2835_GPHEN0 as defined in all/bcm2835.h:837
	BCM2835_GPHEN0 = 100
	// BCM2835_GPHEN1 as defined in all/bcm2835.h:838
	BCM2835_GPHEN1 = 104
	// BCM2835_GPLEN0 as defined in all/bcm2835.h:839
	BCM2835_GPLEN0 = 112
	// BCM2835_GPLEN1 as defined in all/bcm2835.h:840
	BCM2835_GPLEN1 = 116
	// BCM2835_GPAREN0 as defined in all/bcm2835.h:841
	BCM2835_GPAREN0 = 124
	// BCM2835_GPAREN1 as defined in all/bcm2835.h:842
	BCM2835_GPAREN1 = 128
	// BCM2835_GPAFEN0 as defined in all/bcm2835.h:843
	BCM2835_GPAFEN0 = 136
	// BCM2835_GPAFEN1 as defined in all/bcm2835.h:844
	BCM2835_GPAFEN1 = 140
	// BCM2835_GPPUD as defined in all/bcm2835.h:845
	BCM2835_GPPUD = 148
	// BCM2835_GPPUDCLK0 as defined in all/bcm2835.h:846
	BCM2835_GPPUDCLK0 = 152
	// BCM2835_GPPUDCLK1 as defined in all/bcm2835.h:847
	BCM2835_GPPUDCLK1 = 156
	// BCM2835_GPPUPPDN0 as defined in all/bcm2835.h:850
	BCM2835_GPPUPPDN0 = 228
	// BCM2835_GPPUPPDN1 as defined in all/bcm2835.h:851
	BCM2835_GPPUPPDN1 = 232
	// BCM2835_GPPUPPDN2 as defined in all/bcm2835.h:852
	BCM2835_GPPUPPDN2 = 236
	// BCM2835_GPPUPPDN3 as defined in all/bcm2835.h:853
	BCM2835_GPPUPPDN3 = 240
	// BCM2835_GPIO_PUD_ERROR as defined in all/bcm2835.h:882
	BCM2835_GPIO_PUD_ERROR = 8
	// BCM2835_PADS_GPIO_0_27 as defined in all/bcm2835.h:885
	BCM2835_PADS_GPIO_0_27 = 44
	// BCM2835_PADS_GPIO_28_45 as defined in all/bcm2835.h:886
	BCM2835_PADS_GPIO_28_45 = 48
	// BCM2835_PADS_GPIO_46_53 as defined in all/bcm2835.h:887
	BCM2835_PADS_GPIO_46_53 = 52
	// BCM2835_PAD_PASSWRD as defined in all/bcm2835.h:890
	BCM2835_PAD_PASSWRD = 1509949440
	// BCM2835_PAD_SLEW_RATE_UNLIMITED as defined in all/bcm2835.h:891
	BCM2835_PAD_SLEW_RATE_UNLIMITED = 16
	// BCM2835_PAD_HYSTERESIS_ENABLED as defined in all/bcm2835.h:892
	BCM2835_PAD_HYSTERESIS_ENABLED = 8
	// BCM2835_PAD_DRIVE_2mA as defined in all/bcm2835.h:893
	BCM2835_PAD_DRIVE_2mA = 0
	// BCM2835_PAD_DRIVE_4mA as defined in all/bcm2835.h:894
	BCM2835_PAD_DRIVE_4mA = 1
	// BCM2835_PAD_DRIVE_6mA as defined in all/bcm2835.h:895
	BCM2835_PAD_DRIVE_6mA = 2
	// BCM2835_PAD_DRIVE_8mA as defined in all/bcm2835.h:896
	BCM2835_PAD_DRIVE_8mA = 3
	// BCM2835_PAD_DRIVE_10mA as defined in all/bcm2835.h:897
	BCM2835_PAD_DRIVE_10mA = 4
	// BCM2835_PAD_DRIVE_12mA as defined in all/bcm2835.h:898
	BCM2835_PAD_DRIVE_12mA = 5
	// BCM2835_PAD_DRIVE_14mA as defined in all/bcm2835.h:899
	BCM2835_PAD_DRIVE_14mA = 6
	// BCM2835_PAD_DRIVE_16mA as defined in all/bcm2835.h:900
	BCM2835_PAD_DRIVE_16mA = 7
	// BCM2835_AUX_IRQ as defined in all/bcm2835.h:1012
	BCM2835_AUX_IRQ = 0
	// BCM2835_AUX_ENABLE as defined in all/bcm2835.h:1013
	BCM2835_AUX_ENABLE = 4
	// BCM2835_AUX_ENABLE_UART1 as defined in all/bcm2835.h:1015
	BCM2835_AUX_ENABLE_UART1 = 1
	// BCM2835_AUX_ENABLE_SPI0 as defined in all/bcm2835.h:1016
	BCM2835_AUX_ENABLE_SPI0 = 2
	// BCM2835_AUX_ENABLE_SPI1 as defined in all/bcm2835.h:1017
	BCM2835_AUX_ENABLE_SPI1 = 4
	// BCM2835_AUX_SPI_CNTL0 as defined in all/bcm2835.h:1020
	BCM2835_AUX_SPI_CNTL0 = 0
	// BCM2835_AUX_SPI_CNTL1 as defined in all/bcm2835.h:1021
	BCM2835_AUX_SPI_CNTL1 = 4
	// BCM2835_AUX_SPI_STAT as defined in all/bcm2835.h:1022
	BCM2835_AUX_SPI_STAT = 8
	// BCM2835_AUX_SPI_PEEK as defined in all/bcm2835.h:1023
	BCM2835_AUX_SPI_PEEK = 12
	// BCM2835_AUX_SPI_IO as defined in all/bcm2835.h:1024
	BCM2835_AUX_SPI_IO = 32
	// BCM2835_AUX_SPI_TXHOLD as defined in all/bcm2835.h:1025
	BCM2835_AUX_SPI_TXHOLD = 48
	// BCM2835_AUX_SPI_CLOCK_MIN as defined in all/bcm2835.h:1027
	BCM2835_AUX_SPI_CLOCK_MIN = 30500
	// BCM2835_AUX_SPI_CLOCK_MAX as defined in all/bcm2835.h:1028
	BCM2835_AUX_SPI_CLOCK_MAX = 125000000
	// BCM2835_AUX_SPI_CNTL0_SPEED as defined in all/bcm2835.h:1030
	BCM2835_AUX_SPI_CNTL0_SPEED = 4293918720
	// BCM2835_AUX_SPI_CNTL0_SPEED_MAX as defined in all/bcm2835.h:1031
	BCM2835_AUX_SPI_CNTL0_SPEED_MAX = 4095
	// BCM2835_AUX_SPI_CNTL0_SPEED_SHIFT as defined in all/bcm2835.h:1032
	BCM2835_AUX_SPI_CNTL0_SPEED_SHIFT = 20
	// BCM2835_AUX_SPI_CNTL0_CS0_N as defined in all/bcm2835.h:1034
	BCM2835_AUX_SPI_CNTL0_CS0_N = 786432
	// BCM2835_AUX_SPI_CNTL0_CS1_N as defined in all/bcm2835.h:1035
	BCM2835_AUX_SPI_CNTL0_CS1_N = 655360
	// BCM2835_AUX_SPI_CNTL0_CS2_N as defined in all/bcm2835.h:1036
	BCM2835_AUX_SPI_CNTL0_CS2_N = 393216
	// BCM2835_AUX_SPI_CNTL0_POSTINPUT as defined in all/bcm2835.h:1038
	BCM2835_AUX_SPI_CNTL0_POSTINPUT = 65536
	// BCM2835_AUX_SPI_CNTL0_VAR_CS as defined in all/bcm2835.h:1039
	BCM2835_AUX_SPI_CNTL0_VAR_CS = 32768
	// BCM2835_AUX_SPI_CNTL0_VAR_WIDTH as defined in all/bcm2835.h:1040
	BCM2835_AUX_SPI_CNTL0_VAR_WIDTH = 16384
	// BCM2835_AUX_SPI_CNTL0_DOUTHOLD as defined in all/bcm2835.h:1041
	BCM2835_AUX_SPI_CNTL0_DOUTHOLD = 12288
	// BCM2835_AUX_SPI_CNTL0_ENABLE as defined in all/bcm2835.h:1042
	BCM2835_AUX_SPI_CNTL0_ENABLE = 2048
	// BCM2835_AUX_SPI_CNTL0_CPHA_IN as defined in all/bcm2835.h:1043
	BCM2835_AUX_SPI_CNTL0_CPHA_IN = 1024
	// BCM2835_AUX_SPI_CNTL0_CLEARFIFO as defined in all/bcm2835.h:1044
	BCM2835_AUX_SPI_CNTL0_CLEARFIFO = 512
	// BCM2835_AUX_SPI_CNTL0_CPHA_OUT as defined in all/bcm2835.h:1045
	BCM2835_AUX_SPI_CNTL0_CPHA_OUT = 256
	// BCM2835_AUX_SPI_CNTL0_CPOL as defined in all/bcm2835.h:1046
	BCM2835_AUX_SPI_CNTL0_CPOL = 128
	// BCM2835_AUX_SPI_CNTL0_MSBF_OUT as defined in all/bcm2835.h:1047
	BCM2835_AUX_SPI_CNTL0_MSBF_OUT = 64
	// BCM2835_AUX_SPI_CNTL0_SHIFTLEN as defined in all/bcm2835.h:1048
	BCM2835_AUX_SPI_CNTL0_SHIFTLEN = 63
	// BCM2835_AUX_SPI_CNTL1_CSHIGH as defined in all/bcm2835.h:1050
	BCM2835_AUX_SPI_CNTL1_CSHIGH = 1792
	// BCM2835_AUX_SPI_CNTL1_IDLE as defined in all/bcm2835.h:1051
	BCM2835_AUX_SPI_CNTL1_IDLE = 128
	// BCM2835_AUX_SPI_CNTL1_TXEMPTY as defined in all/bcm2835.h:1052
	BCM2835_AUX_SPI_CNTL1_TXEMPTY = 64
	// BCM2835_AUX_SPI_CNTL1_MSBF_IN as defined in all/bcm2835.h:1053
	BCM2835_AUX_SPI_CNTL1_MSBF_IN = 2
	// BCM2835_AUX_SPI_CNTL1_KEEP_IN as defined in all/bcm2835.h:1054
	BCM2835_AUX_SPI_CNTL1_KEEP_IN = 1
	// BCM2835_AUX_SPI_STAT_TX_LVL as defined in all/bcm2835.h:1056
	BCM2835_AUX_SPI_STAT_TX_LVL = 4026531840
	// BCM2835_AUX_SPI_STAT_RX_LVL as defined in all/bcm2835.h:1057
	BCM2835_AUX_SPI_STAT_RX_LVL = 15728640
	// BCM2835_AUX_SPI_STAT_TX_FULL as defined in all/bcm2835.h:1058
	BCM2835_AUX_SPI_STAT_TX_FULL = 1024
	// BCM2835_AUX_SPI_STAT_TX_EMPTY as defined in all/bcm2835.h:1059
	BCM2835_AUX_SPI_STAT_TX_EMPTY = 512
	// BCM2835_AUX_SPI_STAT_RX_FULL as defined in all/bcm2835.h:1060
	BCM2835_AUX_SPI_STAT_RX_FULL = 256
	// BCM2835_AUX_SPI_STAT_RX_EMPTY as defined in all/bcm2835.h:1061
	BCM2835_AUX_SPI_STAT_RX_EMPTY = 128
	// BCM2835_AUX_SPI_STAT_BUSY as defined in all/bcm2835.h:1062
	BCM2835_AUX_SPI_STAT_BUSY = 64
	// BCM2835_AUX_SPI_STAT_BITCOUNT as defined in all/bcm2835.h:1063
	BCM2835_AUX_SPI_STAT_BITCOUNT = 63
	// BCM2835_SPI0_CS as defined in all/bcm2835.h:1069
	BCM2835_SPI0_CS = 0
	// BCM2835_SPI0_FIFO as defined in all/bcm2835.h:1070
	BCM2835_SPI0_FIFO = 4
	// BCM2835_SPI0_CLK as defined in all/bcm2835.h:1071
	BCM2835_SPI0_CLK = 8
	// BCM2835_SPI0_DLEN as defined in all/bcm2835.h:1072
	BCM2835_SPI0_DLEN = 12
	// BCM2835_SPI0_LTOH as defined in all/bcm2835.h:1073
	BCM2835_SPI0_LTOH = 16
	// BCM2835_SPI0_DC as defined in all/bcm2835.h:1074
	BCM2835_SPI0_DC = 20
	// BCM2835_SPI0_CS_LEN_LONG as defined in all/bcm2835.h:1077
	BCM2835_SPI0_CS_LEN_LONG = 33554432
	// BCM2835_SPI0_CS_DMA_LEN as defined in all/bcm2835.h:1078
	BCM2835_SPI0_CS_DMA_LEN = 16777216
	// BCM2835_SPI0_CS_CSPOL2 as defined in all/bcm2835.h:1079
	BCM2835_SPI0_CS_CSPOL2 = 8388608
	// BCM2835_SPI0_CS_CSPOL1 as defined in all/bcm2835.h:1080
	BCM2835_SPI0_CS_CSPOL1 = 4194304
	// BCM2835_SPI0_CS_CSPOL0 as defined in all/bcm2835.h:1081
	BCM2835_SPI0_CS_CSPOL0 = 2097152
	// BCM2835_SPI0_CS_RXF as defined in all/bcm2835.h:1082
	BCM2835_SPI0_CS_RXF = 1048576
	// BCM2835_SPI0_CS_RXR as defined in all/bcm2835.h:1083
	BCM2835_SPI0_CS_RXR = 524288
	// BCM2835_SPI0_CS_TXD as defined in all/bcm2835.h:1084
	BCM2835_SPI0_CS_TXD = 262144
	// BCM2835_SPI0_CS_RXD as defined in all/bcm2835.h:1085
	BCM2835_SPI0_CS_RXD = 131072
	// BCM2835_SPI0_CS_DONE as defined in all/bcm2835.h:1086
	BCM2835_SPI0_CS_DONE = 65536
	// BCM2835_SPI0_CS_TE_EN as defined in all/bcm2835.h:1087
	BCM2835_SPI0_CS_TE_EN = 32768
	// BCM2835_SPI0_CS_LMONO as defined in all/bcm2835.h:1088
	BCM2835_SPI0_CS_LMONO = 16384
	// BCM2835_SPI0_CS_LEN as defined in all/bcm2835.h:1089
	BCM2835_SPI0_CS_LEN = 8192
	// BCM2835_SPI0_CS_REN as defined in all/bcm2835.h:1090
	BCM2835_SPI0_CS_REN = 4096
	// BCM2835_SPI0_CS_ADCS as defined in all/bcm2835.h:1091
	BCM2835_SPI0_CS_ADCS = 2048
	// BCM2835_SPI0_CS_INTR as defined in all/bcm2835.h:1092
	BCM2835_SPI0_CS_INTR = 1024
	// BCM2835_SPI0_CS_INTD as defined in all/bcm2835.h:1093
	BCM2835_SPI0_CS_INTD = 512
	// BCM2835_SPI0_CS_DMAEN as defined in all/bcm2835.h:1094
	BCM2835_SPI0_CS_DMAEN = 256
	// BCM2835_SPI0_CS_TA as defined in all/bcm2835.h:1095
	BCM2835_SPI0_CS_TA = 128
	// BCM2835_SPI0_CS_CSPOL as defined in all/bcm2835.h:1096
	BCM2835_SPI0_CS_CSPOL = 64
	// BCM2835_SPI0_CS_CLEAR as defined in all/bcm2835.h:1097
	BCM2835_SPI0_CS_CLEAR = 48
	// BCM2835_SPI0_CS_CLEAR_RX as defined in all/bcm2835.h:1098
	BCM2835_SPI0_CS_CLEAR_RX = 32
	// BCM2835_SPI0_CS_CLEAR_TX as defined in all/bcm2835.h:1099
	BCM2835_SPI0_CS_CLEAR_TX = 16
	// BCM2835_SPI0_CS_CPOL as defined in all/bcm2835.h:1100
	BCM2835_SPI0_CS_CPOL = 8
	// BCM2835_SPI0_CS_CPHA as defined in all/bcm2835.h:1101
	BCM2835_SPI0_CS_CPHA = 4
	// BCM2835_SPI0_CS_CS as defined in all/bcm2835.h:1102
	BCM2835_SPI0_CS_CS = 3
	// BCM2835_BSC_C as defined in all/bcm2835.h:1174
	BCM2835_BSC_C = 0
	// BCM2835_BSC_S as defined in all/bcm2835.h:1175
	BCM2835_BSC_S = 4
	// BCM2835_BSC_DLEN as defined in all/bcm2835.h:1176
	BCM2835_BSC_DLEN = 8
	// BCM2835_BSC_A as defined in all/bcm2835.h:1177
	BCM2835_BSC_A = 12
	// BCM2835_BSC_FIFO as defined in all/bcm2835.h:1178
	BCM2835_BSC_FIFO = 16
	// BCM2835_BSC_DIV as defined in all/bcm2835.h:1179
	BCM2835_BSC_DIV = 20
	// BCM2835_BSC_DEL as defined in all/bcm2835.h:1180
	BCM2835_BSC_DEL = 24
	// BCM2835_BSC_CLKT as defined in all/bcm2835.h:1181
	BCM2835_BSC_CLKT = 28
	// BCM2835_BSC_C_I2CEN as defined in all/bcm2835.h:1184
	BCM2835_BSC_C_I2CEN = 32768
	// BCM2835_BSC_C_INTR as defined in all/bcm2835.h:1185
	BCM2835_BSC_C_INTR = 1024
	// BCM2835_BSC_C_INTT as defined in all/bcm2835.h:1186
	BCM2835_BSC_C_INTT = 512
	// BCM2835_BSC_C_INTD as defined in all/bcm2835.h:1187
	BCM2835_BSC_C_INTD = 256
	// BCM2835_BSC_C_ST as defined in all/bcm2835.h:1188
	BCM2835_BSC_C_ST = 128
	// BCM2835_BSC_C_CLEAR_1 as defined in all/bcm2835.h:1189
	BCM2835_BSC_C_CLEAR_1 = 32
	// BCM2835_BSC_C_CLEAR_2 as defined in all/bcm2835.h:1190
	BCM2835_BSC_C_CLEAR_2 = 16
	// BCM2835_BSC_C_READ as defined in all/bcm2835.h:1191
	BCM2835_BSC_C_READ = 1
	// BCM2835_BSC_S_CLKT as defined in all/bcm2835.h:1194
	BCM2835_BSC_S_CLKT = 512
	// BCM2835_BSC_S_ERR as defined in all/bcm2835.h:1195
	BCM2835_BSC_S_ERR = 256
	// BCM2835_BSC_S_RXF as defined in all/bcm2835.h:1196
	BCM2835_BSC_S_RXF = 128
	// BCM2835_BSC_S_TXE as defined in all/bcm2835.h:1197
	BCM2835_BSC_S_TXE = 64
	// BCM2835_BSC_S_RXD as defined in all/bcm2835.h:1198
	BCM2835_BSC_S_RXD = 32
	// BCM2835_BSC_S_TXD as defined in all/bcm2835.h:1199
	BCM2835_BSC_S_TXD = 16
	// BCM2835_BSC_S_RXR as defined in all/bcm2835.h:1200
	BCM2835_BSC_S_RXR = 8
	// BCM2835_BSC_S_TXW as defined in all/bcm2835.h:1201
	BCM2835_BSC_S_TXW = 4
	// BCM2835_BSC_S_DONE as defined in all/bcm2835.h:1202
	BCM2835_BSC_S_DONE = 2
	// BCM2835_BSC_S_TA as defined in all/bcm2835.h:1203
	BCM2835_BSC_S_TA = 1
	// BCM2835_BSC_FIFO_SIZE as defined in all/bcm2835.h:1205
	BCM2835_BSC_FIFO_SIZE = 16
	// BCM2835_SMI_CS as defined in all/bcm2835.h:1231
	BCM2835_SMI_CS = 0
	// BCM2835_SMI_LENGTH as defined in all/bcm2835.h:1232
	BCM2835_SMI_LENGTH = 1
	// BCM2835_SMI_ADRS as defined in all/bcm2835.h:1233
	BCM2835_SMI_ADRS = 2
	// BCM2835_SMI_DATA as defined in all/bcm2835.h:1234
	BCM2835_SMI_DATA = 3
	// BCM2835_SMI_READ0 as defined in all/bcm2835.h:1235
	BCM2835_SMI_READ0 = 4
	// BCM2835_SMI_WRITE0 as defined in all/bcm2835.h:1236
	BCM2835_SMI_WRITE0 = 5
	// BCM2835_SMI_READ1 as defined in all/bcm2835.h:1237
	BCM2835_SMI_READ1 = 6
	// BCM2835_SMI_WRITE1 as defined in all/bcm2835.h:1238
	BCM2835_SMI_WRITE1 = 7
	// BCM2835_SMI_READ2 as defined in all/bcm2835.h:1239
	BCM2835_SMI_READ2 = 8
	// BCM2835_SMI_WRITE2 as defined in all/bcm2835.h:1240
	BCM2835_SMI_WRITE2 = 9
	// BCM2835_SMI_READ3 as defined in all/bcm2835.h:1241
	BCM2835_SMI_READ3 = 10
	// BCM2835_SMI_WRITE3 as defined in all/bcm2835.h:1242
	BCM2835_SMI_WRITE3 = 11
	// BCM2835_SMI_DMAC as defined in all/bcm2835.h:1243
	BCM2835_SMI_DMAC = 12
	// BCM2835_SMI_DIRCS as defined in all/bcm2835.h:1244
	BCM2835_SMI_DIRCS = 13
	// BCM2835_SMI_DIRADDR as defined in all/bcm2835.h:1245
	BCM2835_SMI_DIRADDR = 14
	// BCM2835_SMI_DIRDATA as defined in all/bcm2835.h:1246
	BCM2835_SMI_DIRDATA = 15
	// BCM2835_SMI_RW_WIDTH_MSK as defined in all/bcm2835.h:1249
	BCM2835_SMI_RW_WIDTH_MSK = 3221225472
	// BCM2835_SMI_RW_WID8 as defined in all/bcm2835.h:1250
	BCM2835_SMI_RW_WID8 = 0
	// BCM2835_SMI_RW_WID16 as defined in all/bcm2835.h:1251
	BCM2835_SMI_RW_WID16 = 1073741824
	// BCM2835_SMI_RW_WID18 as defined in all/bcm2835.h:1252
	BCM2835_SMI_RW_WID18 = 2147483648
	// BCM2835_SMI_RW_WID9 as defined in all/bcm2835.h:1253
	BCM2835_SMI_RW_WID9 = 3221225472
	// BCM2835_SMI_RW_SETUP_MSK as defined in all/bcm2835.h:1254
	BCM2835_SMI_RW_SETUP_MSK = 1056964608
	// BCM2835_SMI_RW_SETUP_LS as defined in all/bcm2835.h:1255
	BCM2835_SMI_RW_SETUP_LS = 24
	// BCM2835_SMI_RW_MODE68 as defined in all/bcm2835.h:1256
	BCM2835_SMI_RW_MODE68 = 8388608
	// BCM2835_SMI_RW_MODE80 as defined in all/bcm2835.h:1257
	BCM2835_SMI_RW_MODE80 = 0
	// BCM2835_SMI_READ_FSETUP as defined in all/bcm2835.h:1258
	BCM2835_SMI_READ_FSETUP = 4194304
	// BCM2835_SMI_WRITE_SWAP as defined in all/bcm2835.h:1259
	BCM2835_SMI_WRITE_SWAP = 4194304
	// BCM2835_SMI_RW_HOLD_MSK as defined in all/bcm2835.h:1260
	BCM2835_SMI_RW_HOLD_MSK = 4128768
	// BCM2835_SMI_RW_HOLD_LS as defined in all/bcm2835.h:1261
	BCM2835_SMI_RW_HOLD_LS = 16
	// BCM2835_SMI_RW_PACEALL as defined in all/bcm2835.h:1262
	BCM2835_SMI_RW_PACEALL = 32768
	// BCM2835_SMI_RW_PACE_MSK as defined in all/bcm2835.h:1263
	BCM2835_SMI_RW_PACE_MSK = 32512
	// BCM2835_SMI_RW_PACE_LS as defined in all/bcm2835.h:1264
	BCM2835_SMI_RW_PACE_LS = 8
	// BCM2835_SMI_RW_DREQ as defined in all/bcm2835.h:1265
	BCM2835_SMI_RW_DREQ = 128
	// BCM2835_SMI_RW_STROBE_MSK as defined in all/bcm2835.h:1266
	BCM2835_SMI_RW_STROBE_MSK = 127
	// BCM2835_SMI_RW_STROBE_LS as defined in all/bcm2835.h:1267
	BCM2835_SMI_RW_STROBE_LS = 0
	// BCM2835_SMI_DIRCS_ENABLE as defined in all/bcm2835.h:1270
	BCM2835_SMI_DIRCS_ENABLE = 1
	// BCM2835_SMI_DIRCS_START as defined in all/bcm2835.h:1271
	BCM2835_SMI_DIRCS_START = 2
	// BCM2835_SMI_DIRCS_DONE as defined in all/bcm2835.h:1272
	BCM2835_SMI_DIRCS_DONE = 4
	// BCM2835_SMI_DIRCS_WRITE as defined in all/bcm2835.h:1273
	BCM2835_SMI_DIRCS_WRITE = 8
	// BCM2835_SMI_DIRADRS_DEV_MSK as defined in all/bcm2835.h:1276
	BCM2835_SMI_DIRADRS_DEV_MSK = 768
	// BCM2835_SMI_DIRADRS_DEV_LS as defined in all/bcm2835.h:1277
	BCM2835_SMI_DIRADRS_DEV_LS = 8
	// BCM2835_SMI_DIRADRS_DEV0 as defined in all/bcm2835.h:1278
	BCM2835_SMI_DIRADRS_DEV0 = 0
	// BCM2835_SMI_DIRADRS_DEV1 as defined in all/bcm2835.h:1279
	BCM2835_SMI_DIRADRS_DEV1 = 256
	// BCM2835_SMI_DIRADRS_DEV2 as defined in all/bcm2835.h:1280
	BCM2835_SMI_DIRADRS_DEV2 = 512
	// BCM2835_SMI_DIRADRS_DEV3 as defined in all/bcm2835.h:1281
	BCM2835_SMI_DIRADRS_DEV3 = 768
	// BCM2835_SMI_DIRADRS_MSK as defined in all/bcm2835.h:1282
	BCM2835_SMI_DIRADRS_MSK = 63
	// BCM2835_SMI_DIRADRS_LS as defined in all/bcm2835.h:1283
	BCM2835_SMI_DIRADRS_LS = 0
	// BCM2835_ST_CS as defined in all/bcm2835.h:1301
	BCM2835_ST_CS = 0
	// BCM2835_ST_CLO as defined in all/bcm2835.h:1302
	BCM2835_ST_CLO = 4
	// BCM2835_ST_CHI as defined in all/bcm2835.h:1303
	BCM2835_ST_CHI = 8
	// BCM2835_PWM_CONTROL as defined in all/bcm2835.h:1309
	BCM2835_PWM_CONTROL = 0
	// BCM2835_PWM_STATUS as defined in all/bcm2835.h:1310
	BCM2835_PWM_STATUS = 1
	// BCM2835_PWM_DMAC as defined in all/bcm2835.h:1311
	BCM2835_PWM_DMAC = 2
	// BCM2835_PWM0_RANGE as defined in all/bcm2835.h:1312
	BCM2835_PWM0_RANGE = 4
	// BCM2835_PWM0_DATA as defined in all/bcm2835.h:1313
	BCM2835_PWM0_DATA = 5
	// BCM2835_PWM_FIF1 as defined in all/bcm2835.h:1314
	BCM2835_PWM_FIF1 = 6
	// BCM2835_PWM1_RANGE as defined in all/bcm2835.h:1315
	BCM2835_PWM1_RANGE = 8
	// BCM2835_PWM1_DATA as defined in all/bcm2835.h:1316
	BCM2835_PWM1_DATA = 9
	// BCM2835_PWMCLK_CNTL as defined in all/bcm2835.h:1319
	BCM2835_PWMCLK_CNTL = 40
	// BCM2835_PWMCLK_DIV as defined in all/bcm2835.h:1320
	BCM2835_PWMCLK_DIV = 41
	// BCM2835_PWM_PASSWRD as defined in all/bcm2835.h:1321
	BCM2835_PWM_PASSWRD = 1509949440
	// BCM2835_PWM1_MS_MODE as defined in all/bcm2835.h:1323
	BCM2835_PWM1_MS_MODE = 32768
	// BCM2835_PWM1_USEFIFO as defined in all/bcm2835.h:1324
	BCM2835_PWM1_USEFIFO = 8192
	// BCM2835_PWM1_REVPOLAR as defined in all/bcm2835.h:1325
	BCM2835_PWM1_REVPOLAR = 4096
	// BCM2835_PWM1_OFFSTATE as defined in all/bcm2835.h:1326
	BCM2835_PWM1_OFFSTATE = 2048
	// BCM2835_PWM1_REPEATFF as defined in all/bcm2835.h:1327
	BCM2835_PWM1_REPEATFF = 1024
	// BCM2835_PWM1_SERIAL as defined in all/bcm2835.h:1328
	BCM2835_PWM1_SERIAL = 512
	// BCM2835_PWM1_ENABLE as defined in all/bcm2835.h:1329
	BCM2835_PWM1_ENABLE = 256
	// BCM2835_PWM0_MS_MODE as defined in all/bcm2835.h:1331
	BCM2835_PWM0_MS_MODE = 128
	// BCM2835_PWM_CLEAR_FIFO as defined in all/bcm2835.h:1332
	BCM2835_PWM_CLEAR_FIFO = 64
	// BCM2835_PWM0_USEFIFO as defined in all/bcm2835.h:1333
	BCM2835_PWM0_USEFIFO = 32
	// BCM2835_PWM0_REVPOLAR as defined in all/bcm2835.h:1334
	BCM2835_PWM0_REVPOLAR = 16
	// BCM2835_PWM0_OFFSTATE as defined in all/bcm2835.h:1335
	BCM2835_PWM0_OFFSTATE = 8
	// BCM2835_PWM0_REPEATFF as defined in all/bcm2835.h:1336
	BCM2835_PWM0_REPEATFF = 4
	// BCM2835_PWM0_SERIAL as defined in all/bcm2835.h:1337
	BCM2835_PWM0_SERIAL = 2
	// BCM2835_PWM0_ENABLE as defined in all/bcm2835.h:1338
	BCM2835_PWM0_ENABLE = 1
	// IT8951_TCON_SYS_RUN as defined in all/EPD_IT8951.h:79
	IT8951_TCON_SYS_RUN = 1
	// IT8951_TCON_STANDBY as defined in all/EPD_IT8951.h:80
	IT8951_TCON_STANDBY = 2
	// IT8951_TCON_SLEEP as defined in all/EPD_IT8951.h:81
	IT8951_TCON_SLEEP = 3
	// IT8951_TCON_REG_RD as defined in all/EPD_IT8951.h:82
	IT8951_TCON_REG_RD = 16
	// IT8951_TCON_REG_WR as defined in all/EPD_IT8951.h:83
	IT8951_TCON_REG_WR = 17
	// IT8951_TCON_MEM_BST_RD_T as defined in all/EPD_IT8951.h:85
	IT8951_TCON_MEM_BST_RD_T = 18
	// IT8951_TCON_MEM_BST_RD_S as defined in all/EPD_IT8951.h:86
	IT8951_TCON_MEM_BST_RD_S = 19
	// IT8951_TCON_MEM_BST_WR as defined in all/EPD_IT8951.h:87
	IT8951_TCON_MEM_BST_WR = 20
	// IT8951_TCON_MEM_BST_END as defined in all/EPD_IT8951.h:88
	IT8951_TCON_MEM_BST_END = 21
	// IT8951_TCON_LD_IMG as defined in all/EPD_IT8951.h:90
	IT8951_TCON_LD_IMG = 32
	// IT8951_TCON_LD_IMG_AREA as defined in all/EPD_IT8951.h:91
	IT8951_TCON_LD_IMG_AREA = 33
	// IT8951_TCON_LD_IMG_END as defined in all/EPD_IT8951.h:92
	IT8951_TCON_LD_IMG_END = 34
	// IT8951_ROTATE_0 as defined in all/EPD_IT8951.h:105
	IT8951_ROTATE_0 = 0
	// IT8951_ROTATE_90 as defined in all/EPD_IT8951.h:106
	IT8951_ROTATE_90 = 1
	// IT8951_ROTATE_180 as defined in all/EPD_IT8951.h:107
	IT8951_ROTATE_180 = 2
	// IT8951_ROTATE_270 as defined in all/EPD_IT8951.h:108
	IT8951_ROTATE_270 = 3
	// IT8951_2BPP as defined in all/EPD_IT8951.h:111
	IT8951_2BPP = 0
	// IT8951_3BPP as defined in all/EPD_IT8951.h:112
	IT8951_3BPP = 1
	// IT8951_4BPP as defined in all/EPD_IT8951.h:113
	IT8951_4BPP = 2
	// IT8951_8BPP as defined in all/EPD_IT8951.h:114
	IT8951_8BPP = 3
	// IT8951_LDIMG_L_ENDIAN as defined in all/EPD_IT8951.h:117
	IT8951_LDIMG_L_ENDIAN = 0
	// IT8951_LDIMG_B_ENDIAN as defined in all/EPD_IT8951.h:118
	IT8951_LDIMG_B_ENDIAN = 1
)

const ()

const ()

const ()