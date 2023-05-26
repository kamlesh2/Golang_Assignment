CREATE MIGRATION m1wdrbf2y7gr6kdz5rgitxdxltufpe7oh4742oxnd3mmmwv22n65dq
    ONTO m16smqiffrfejvy3augh3t5gv6entzyigx3h46sbvcsmsajdtgc6cq
{
  CREATE TYPE default::Resume {
      CREATE REQUIRED PROPERTY fname -> std::str;
      CREATE REQUIRED PROPERTY lname -> std::str;
  };
};
