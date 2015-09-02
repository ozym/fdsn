# FDSN

A golang wrapper package for the __FDSN__ _StationXML_ format for describing seismic data collection parameters.

[![Build Status](https://travis-ci.org/ozym/fdsn.svg?branch=master)](https://travis-ci.org/ozym/fdsn)

See http://www.fdsn.org/xml/station/ for schema details.

Currently the classes are capable of marshalling / unmarshalling a set of test _StationXML_ files derived from the NZ
__FDSN__ webservice.

Expected enhancements:

* Validation
* Enumeration of selections
* More tests

Progress:

- [x]RootType
- [ ]NetworkType
- [ ]StationType
- [ ]ChannelType
- [ ]GainType
- [ ]FrequencyRangeGroup
- [ ]SensitivityType
- [ ]EquipmentType
- [ ]ResponseStageType
- [ ]LogType
- [ ]CommentType
- [ ]PolesZerosType
- [ ]FIRType
- [ ]CoefficientsType
- [ ]ResponseListElementType
- [ ]ResponseListType
- [ ]PolynomialType
- [ ]DecimationType
- [ ]uncertaintyDouble
- [ ]FloatNoUnitType
- [ ]FloatType
- [ ]SecondType
- [ ]VoltageType
- [ ]AngleType
- [ ]LatitudeBaseType
- [ ]LatitudeType
- [ ]LongitudeBaseType
- [ ]LongitudeType
- [ ]AzimuthType
- [ ]DipType
- [ ]DistanceType
- [ ]FrequencyType
- [ ]RateGroup
- [ ]SampleRateType
- [ ]SampleRateRatioType
- [ ]PoleZeroType
- [ ]CounterType
- [ ]PersonType
- [ ]SiteType
- [x]ExternalReferenceType
- [ ]NominalType
- [ ]EmailType
- [ ]PhoneNumberType
- [ ]RestrictedStatusType
- [ ]UnitsType
- [ ]BaseFilterType
- [ ]ResponseType
- [ ]BaseNodeType
