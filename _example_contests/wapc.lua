-- Basic functions for getting contest metadata
function metadata() {
    local meta = {
        -- basic metadata of the contest
        -- meta_version: the schema version of the metadata
        meta_version: 1
        -- version: the version of the contest, for display only
        version: "2022-10-02"

        -- the name of the contest, for display only
        display_name: "WAPC CW"

        -- the sent exchange data of the contest, for display only
        exchange_sent: {"RST", "Province"}
        -- the received exchange data of the contest, for display only
        exchange_recv: {"RST", "Province/Seq No."}

        -- multipliers are only used for display, score calculation is fully handled in the script
        multipliers: {"Province", "DXCC"}

        -- The fields in the "exchange_sent" which can be set before the contest
        customize_fields: {"Province"}
    }

    return meta
}