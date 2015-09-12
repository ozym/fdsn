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

Progress: String() IsValid()

- [x] RootType
- [x] NetworkType
- [ ] StationType
- [ ] ChannelType
- [x] GainType
- [ ] FrequencyRangeGroup
- [x] SensitivityType
- [ ] EquipmentType
- [ ] ResponseStageType
- [ ] LogType
- [ ] CommentType
- [ ] PolesZerosType
- [x] FIRType
- [ ] CoefficientsType
- [ ] ResponseListElementType
- [ ] ResponseListType
- [ ] PolynomialType
- [ ] DecimationType
- [ ] uncertaintyDouble
- [ ] FloatNoUnitType
- [ ] FloatType
- [ ] SecondType
- [ ] VoltageType
- [ ] AngleType
- [x] LatitudeBaseType
- [x] LatitudeType
- [x] LongitudeBaseType
- [x] LongitudeType
- [ ] AzimuthType
- [ ] DipType
- [x] DistanceType
- [ ] FrequencyType
- [x] SampleRateGroup
- [x] SampleRateType
- [x] SampleRateRatioType
- [x] PoleZeroType
- [x] CounterType
- [x] PersonType
- [x] SiteType
- [x] ExternalReferenceType
- [ ] NominalType
- [x] EmailType
- [x] PhoneNumberType
- [x] RestrictedStatusType
- [x] UnitsType
- [x] BaseFilterType
- [x] ResponseType
- [x] BaseNodeType
