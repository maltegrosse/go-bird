package bird

import (
	"errors"
	"math"
)

///////////////////////////////////////////////
//             Richard E. Bird               //
// Clear Sky Broadband Solar Radiation Model //
//                                           //
//            September 19, 2012             //
//                                           //
//                                           //
//   Afshin Michael Andreas                  //
//   Afshin.Andreas@NREL.gov (303)384-6383   //
//                                           //
//   Solar Resource and Forecasting Group    //
//   Solar Radiation Research Laboratory     //
//   National Renewable Energy Laboratory    //
//   15013 Denver W Pkwy, Golden, CO 80401   //
///////////////////////////////////////////////

///////////////////////////////////////////////
//  This code is based on the SERI (NREL)    //
//  technical report "A Simplified Clear	 //
//  Sky model for Direct and Diffuse 		 //
//  Insolation on Horizontal Surfaces" by    //
//  R.E. Bird and R.L. Hulstrom              //
///////////////////////////////////////////////

// Bird interface defines the public functions
type Bird interface {
	/* INPUT */
	// validates and calculates
	Calculate() error
	//solar zenith angle [degrees] -- available from SPA output
	GetZenith() float64
	SetZenith(zenith float64)
	//earth radius vector [Astronomical Units, AU] -- available from SPA output
	GetR() float64
	SetR(r float64)
	//annual average local pressure [millibars] -- available from SPA input
	GetPressure() float64
	SetPressure(pressure float64)
	//total column ozone thickness [cm] -- range from 0.05 - 0.4
	GetOzone() float64
	SetOzone(ozone float64)
	//total column water vapor [cm] -- range from 0.01 - 6.5
	GetWater() float64
	SetWater(water float64)
	//broadband aerosol optical depth -- range from 0.02 - 0.5
	GetTaua() float64
	SetTaua(taua float64)
	//forward scattering factor -- 0.85 recommended for rural aerosols
	GetBa() float64
	SetBa(ba float64)
	//ground reflectance -- earth typical is 0.2, snow 0.9, vegetation 0.25
	GetAlbedo() float64
	SetAlbedo(albedo float64)
	//direct normal irradiance modification factor -- optional value from 0.0 - 1.0, which is used to calculate the second set of "modified" irradiance values
	GetDniMod() float64
	SetDniMod(dniMod float64)

	/* OUTPUT */
	//relative optical airmass (not pressure corrected)
	GetAmass() float64
	//direct normal solar irradiance [W/m^2] -- Bird Clear Sky Estimated
	GetDirectNormal() float64
	//global horizontal solar irradiance [W/m^2] -- Bird Clear Sky Estimated
	GetGlobalHoriz() float64
	//diffuse horizontal solar irradiance [W/m^2] -- Bird Clear Sky Estimated
	GetDiffuseHoriz() float64

	//equavalent to direct_normal * dni_mod
	GetDirectNormalMod() float64
	//re-computed global horizontal based on direct_normal_mod
	GetGlobalHorizMod() float64
	//re-computed diffuse horizontal based on direct_normal_mod
	GetDiffuseHorizMod() float64
}

// NewBird creates new Bird instance
func NewBird(zenith float64, r float64, pressure float64, ozone float64, water float64, taua float64, ba float64, albedo float64, dniMod float64) (Bird, error) {
	var b bird
	b.zenith = zenith
	b.r = r
	b.pressure = pressure
	b.ozone = ozone
	b.water = water
	b.taua = taua
	b.ba = ba
	b.albedo = albedo
	b.dniMod = dniMod

	return &b, b.Calculate()
}

type bird struct {
	//----------------------------------INPUT VALUES--------------------------------------------
	zenith   float64 //solar zenith angle [degrees] -- available from SPA output
	r        float64 //earth radius vector [Astronomical Units, AU] -- available from SPA output
	pressure float64 //annual average local pressure [millibars] -- available from SPA input

	ozone  float64 //total column ozone thickness [cm] -- range from 0.05 - 0.4
	water  float64 //total column water vapor [cm] -- range from 0.01 - 6.5
	taua   float64 //broadband aerosol optical depth -- range from 0.02 - 0.5
	ba     float64 //forward scattering factor -- 0.85 recommended for rural aerosols
	albedo float64 //ground reflectance -- earth typical is 0.2, snow 0.9, vegetation 0.25
	dniMod float64 //direct normal irradiance modification factor -- optional value from 0.0 - 1.0,
	//   which is used to calculate the second set of "modified" irradiance values

	//--------------------------------- OUTPUT VALUES-------------------------------------------
	amass float64 //relative optical airmass (not pressure corrected)

	directNormal float64 //direct normal solar irradiance [W/m^2] -- Bird Clear Sky Estimated
	globalHoriz  float64 //global horizontal solar irradiance [W/m^2] -- Bird Clear Sky Estimated
	diffuseHoriz float64 //diffuse horizontal solar irradiance [W/m^2] -- Bird Clear Sky Estimated

	directNormalMod float64 //equavalent to direct_normal * dni_mod
	globalHorizMod  float64 //re-computed global horizontal based on direct_normal_mod
	diffuseHorizMod float64 //re-computed diffuse horizontal based on direct_normal_mod
}

func (b *bird) GetZenith() float64 {
	return b.zenith
}

func (b *bird) SetZenith(zenith float64) {
	b.zenith = zenith
}

func (b *bird) GetR() float64 {
	return b.r
}

func (b *bird) SetR(r float64) {
	b.r = r
}

func (b *bird) GetPressure() float64 {
	return b.pressure
}

func (b *bird) SetPressure(pressure float64) {
	b.pressure = pressure
}

func (b *bird) GetOzone() float64 {
	return b.ozone
}

func (b *bird) SetOzone(ozone float64) {
	b.ozone = ozone
}

func (b *bird) GetWater() float64 {
	return b.water
}

func (b *bird) SetWater(water float64) {
	b.water = water
}

func (b *bird) GetTaua() float64 {
	return b.taua
}

func (b *bird) SetTaua(taua float64) {
	b.taua = taua
}

func (b *bird) GetBa() float64 {
	return b.ba
}

func (b *bird) SetBa(ba float64) {
	b.ba = ba
}

func (b *bird) GetAlbedo() float64 {
	return b.albedo
}

func (b *bird) SetAlbedo(albedo float64) {
	b.albedo = albedo
}

func (b *bird) GetDniMod() float64 {
	return b.dniMod
}

func (b *bird) SetDniMod(dniMod float64) {
	b.dniMod = dniMod
}

func (b *bird) GetAmass() float64 {
	return b.amass
}

func (b *bird) GetDirectNormal() float64 {
	return b.directNormal
}

func (b *bird) GetGlobalHoriz() float64 {
	return b.globalHoriz
}

func (b *bird) GetDiffuseHoriz() float64 {
	return b.diffuseHoriz
}

func (b *bird) GetDirectNormalMod() float64 {
	return b.directNormalMod
}

func (b *bird) GetGlobalHorizMod() float64 {
	return b.globalHorizMod
}

func (b *bird) GetDiffuseHorizMod() float64 {
	return b.diffuseHorizMod
}

func (b *bird) Calculate() error {
	err := b.validate()
	if err != nil {
		return err
	}
	var etrn, press, oz, wat, coszen, taa, rs, ias float64
	var tRayliegh, tOzone, tGases, tWater, tAerosol float64

	if (b.zenith >= 0) && (b.zenith < 90) && (b.r > 0) {
		etrn = 1367.0 / (b.r * b.r)

		coszen = math.Cos(b.zenith * math.Pi / 180.0)
		b.amass = 1.0 / (coszen + 0.50572*math.Pow(96.07995-b.zenith, -1.6364))

		press = b.pressure * b.amass / 1013
		oz = b.ozone * b.amass
		wat = b.water * b.amass

		tRayliegh = math.Exp(-0.0903 * math.Pow(press, 0.84) * (1 + press - math.Pow(press, 1.01)))
		tOzone = 1 - 0.1611*(oz)*math.Pow(1+139.48*oz, -0.3034) - 0.002715*(oz)/(1+0.044*(oz)+0.0003*oz*oz)
		tGases = math.Exp(-0.0127 * math.Pow(press, 0.26))
		tWater = 1 - 2.4959*wat/(math.Pow(1+79.034*wat, 0.6828)+6.385*wat)
		tAerosol = math.Exp(-(math.Pow(b.taua, 0.873)) * (1 + b.taua - math.Pow(b.taua, 0.7088)) * math.Pow(b.amass, 0.9108))

		b.directNormal = 0.9662 * etrn * tAerosol * tWater * tGases * tOzone * tRayliegh

		taa = 1 - 0.1*(1-b.amass+math.Pow(b.amass, 1.06))*(1-tAerosol)
		rs = 0.0685 + (1-b.ba)*(1-tAerosol/taa)
		ias = etrn * coszen * 0.79 * tOzone * tGases * tWater * taa * (0.5*(1-tRayliegh) + b.ba*(1-(tAerosol/taa))) / (1 - b.amass + math.Pow(b.amass, 1.02))

		b.birdHoriz(b.directNormal, coszen, ias, b.albedo, rs, &b.globalHoriz, &b.diffuseHoriz)

		if b.dniMod >= 0 {
			b.directNormalMod = b.directNormal * b.dniMod
			b.birdHoriz(b.directNormalMod, coszen, ias, b.albedo, rs, &b.globalHorizMod, &b.diffuseHorizMod)
		}

	} else {
		b.amass, b.directNormal, b.globalHoriz, b.diffuseHoriz = 0, 0, 0, 0
		b.directNormalMod, b.globalHorizMod, b.diffuseHorizMod = 0, 0, 0
	}
	return nil
}

func (b *bird) birdHoriz(dni float64, coszen float64, ias float64, albedo float64, rs float64, ghi *float64, dhi *float64) {
	dirHorz := dni * coszen

	*ghi = (dirHorz + ias) / (1 - albedo*rs)
	*dhi = *ghi - dirHorz
}

func (b *bird) validate() error {
	if (b.ozone < 0) || (b.ozone > 100) {
		return errors.New("invalid ozone thickness [cm]")
	}
	if (b.water < 0) || (b.water > 100) {
		return errors.New("invalid water vapor [cm]")
	}
	if (b.taua < 0) || (b.taua > 100) {
		return errors.New("invalid broadband aerosol optical depth")
	}
	if (b.ba < 0) || (b.ba > 100) {
		return errors.New("invalid forward scattering factor")
	}
	if (b.albedo < 0) || (b.albedo > 100) {
		return errors.New("invalid ground reflectance")
	}
	if (b.dniMod < 0) || (b.dniMod > 100) {
		return errors.New("invalid direct normal irradiance modification factor")
	}
	return nil
}
