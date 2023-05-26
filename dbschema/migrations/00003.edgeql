CREATE MIGRATION m1j5ke3qmu2tinqcz4jryafkm4bahjzaic3nzcv2ysacaic6xhhgaa
    ONTO m1wdrbf2y7gr6kdz5rgitxdxltufpe7oh4742oxnd3mmmwv22n65dq
{
  ALTER TYPE default::Resume {
      CREATE PROPERTY dob -> std::str;
  };
  ALTER TYPE default::Resume {
      CREATE PROPERTY email -> std::str;
  };
  ALTER TYPE default::Resume {
      ALTER PROPERTY fname {
          RESET OPTIONALITY;
      };
  };
  ALTER TYPE default::Resume {
      ALTER PROPERTY lname {
          RESET OPTIONALITY;
      };
  };
  ALTER TYPE default::Resume {
      CREATE PROPERTY phone -> std::str;
  };
};
