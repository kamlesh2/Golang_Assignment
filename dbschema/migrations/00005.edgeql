CREATE MIGRATION m1zjgwvdubdid2gxjg3ujnurpp2y3yeq6xry6kqv7borvmnqhfhj6a
    ONTO m1k2ty4i4pnofomfm5ajd6fhm24mugbq5kebgzktz7jaxz62ms2cpa
{
  ALTER TYPE default::Resume {
      ALTER PROPERTY fname {
          SET REQUIRED USING ('test');
      };
  };
};
