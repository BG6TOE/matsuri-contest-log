Context Definition Script Spec
==============================

*Current version: 1*

0. The scripts should be written in Lua language.
0. Functions
   1.  `metadata()` -- return the metadata of the contest. Should be a table of the following fields:
       - `api_version` The API version of the metadata, should be `1` for now.
       - `version` A string of the version of the script, can be any string.
       - `display_name` A string for the name of the contest, can be any string.
       - `available_modes` A table of modes available in the contest. Suggested to be a subset of `["SSB", "CW", "DATA"]`, other available values are `["LSB", "USB", "CW", "FT8", "FT4", "JT65", "PSK", "PSK31", "RTTY", "FM"]` and `["PHONE"]`, or even `["ANY"]`.
          > Note: This will be a reference for checking multipliers, you should use a the mode used when calculating the score. (i.e. Use "LSB" and "USB" only when they are considered as two different modes when calculating score, and use `ANY` if the contest does not care about the detailed modes.)
       - `available_bands` A table of bands available in the contest. Must be a subset of `["1600", "160", "80", "60", "40", "30", "20", "17", "15", "12", "10", "6", "2", "70cm", "1G2", "2G4", "5G6", "10G", ...]`
       - `exchange_sent` A table of fields of the exchange sent.
         > Note: Special names of the field: `_rst`, used to indicate the RST sent, which will be default to "599" for cw, "59" for phone and "595" for data modes. You can use `RST` to indicate that the users are suggested to use real signal reports during the contest.
       - `exchange_rcvd` A table of fields of the exchange received.
         > Note: Special names of the field: `_rst`, used to indicate the RST received, which will be default to "599" for cw, "59" for phone and "595" for data modes. You can use `RST` to indicate that the users are suggested to use real signal reports during the contest.
       - `multipliers` Declared multipliers of the contest, display use only.
       - `const_fields` Table of one or more field names in `exchange_sent`, for users to set it as a part of contest parameter. For example, the age of the operator can be included here.
       - `field_description` A brief description of some fields.