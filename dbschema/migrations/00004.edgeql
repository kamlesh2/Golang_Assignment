CREATE MIGRATION m1k2ty4i4pnofomfm5ajd6fhm24mugbq5kebgzktz7jaxz62ms2cpa
    ONTO m1j5ke3qmu2tinqcz4jryafkm4bahjzaic3nzcv2ysacaic6xhhgaa
{
  ALTER TYPE default::Resume {
      CREATE PROPERTY file -> std::str;
  };
};
