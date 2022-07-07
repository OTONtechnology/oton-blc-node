package tx



const (
 TypeOK uint32 = 0
 TypeEncodingError uint32 = 1
 TypeBadNonce uint32 = 2
 TypeUnauthorized uint32 = 3
 TypeUnknownError uint32 = 4


 TypeErrInput uint32 = 5
 TypeErrOutput uint32 = 6
 TypeErrApp uint32 = 7
 TypeErrTxSize uint32 = 8
 TypeErrTxDecode uint32 = 9
 TypeErrDeliverTx uint32 = 10
 TypeErrCheckTx uint32 = 11


 CTErrValidateInputsBasic uint32 = 12
 CTErrValidateOutputsBasic uint32 = 14
 CTErrGetInputs uint32 = 15
 CTErrInputsNoMovable uint32 = 151
 CTErrGetOrMakeOutputs uint32 = 16

 CTErrValidateInputsAdvanced uint32 = 17
 ErrInvalidSequence uint32 = 170
 ErrInputsBalance uint32 = 171
 ErrInvalidSignature uint32 = 172

 CTErrBaseInvalidOutput uint32 = 18
 CTErrInputsCoinsNotEqualFee uint32 = 19
 CTErrValidateNewCoins uint32 = 20
 TypeErrMinterOuner uint32 = 201
 TypeErrAgeOuner uint32 = 301
 CTErrAdjustCoins uint32 = 21


 TypeErrNewAMC uint32 = 22
 TypeErrDubAMC uint32 = 23
 TypeErrParamAMC uint32 = 24

 TypeErrGetAMC uint32 = 25
 TypeErrSetAMC uint32 = 26

 TypeErrMasterAMC uint32 = 28
 TypeErrSponsorAMC uint32 = 29
 TypeErrReferalAMC uint32 = 291
 TypeErrTreeAMC uint32 = 281
 TypeErrRefChain uint32 = 282
 TypeErrBuyAMC uint32 = 283


 TypeErrOldAddress uint32 = 30
 TypeErrWantList uint32 = 31
 TypeErrChangeAdr uint32 = 32
 TypeErrVoteAdr uint32 = 33


 TypeErrValOuner uint32 = 40
)
