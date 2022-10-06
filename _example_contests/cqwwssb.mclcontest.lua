/**
 * Note: The contest descriptor should not use any global variables.
 * Functions will be provided with a "state" for saving the state
 */

/**
 * MCL runtime provides the following lookup function:
 * callsign_cq_zone()   -- Return the inferred CQ zone of the callsign, 0 for unknown
 * callsign_itu_zone()  -- Return the inferred ITU zone of the callsign, 0 for unknown
 * callsign_continent() -- Return the inferred continent of the callsign, "" for unknown
 * callsign_dxcc()      -- Return the inferred DXCC of the callsign, 0 for unknown
 */

-- Basic functions for getting contest metadata
function metadata()
    local meta = {
        -- basic metadata of the contest
        -- meta_version: the schema version of the metadata
        api_version = 1,
        -- version: the version of the contest, for display only
        version = "2022-10-02",

        -- the name of the contest, for display only
        display_name = "CQWW SSB",

        -- the modes of the contest
        available_modes = {"SSB"},

        -- the bands of the contest
        available_bands = {"160", "80", "40", "20", "15", "10"},

        -- the sent exchange data of the contest, for display only
        exchange_sent = {"_rst", "Zone"},
        -- the received exchange data of the contest, for display only
        exchange_rcvd = {"_rst", "Zone"},

        -- multipliers are only used for display, score calculation is fully handled in the script
        multipliers = {"Zone"},

        -- The fields in the "exchange_sent" which can be set before the contest
        const_fields = {"Zone"},

        -- Explain some fields to the user.
        field_description = { Zone = "CQ Zone", }
    }

    return meta
end

-- Create a new state of the contest
-- MCL will not inteprete the content of the state, it is totally private.
function init_state(station)
    return {
        zone_band = {160 = {}, 80 = {}, 40 = {}, 20 = {}, 15 = {}, 10 = {}},
        qso_band = {160 = {}, 80 = {}, 40 = {}, 20 = {}, 15 = {}, 10 = {}},
        qsos = {160 = {}, 80 = {}, 40 = {}, 20 = {}, 15 = {}, 10 = {}},
    }
end

-- TODO QSO validator
function valid_exchange(exch)
    return true
end

-- TODO Get next contest date from a given day

-- TODO Score Calculator
function add_qso(state, qso)
    local score = {
        title = {"Band", "QSO", "Zone", "Score", "Score/Q"},
        fields = {
            {"80", 0, 0, 0, 0},
        },
    }
    return state, score
end

-- TODO adif exporter
-- map exchange sent & exchange rcvd to ADIF field name
function adif_mapping()
    return {}
end

-- TODO check callsign, called when trying to input a new QSO
-- Should not update the state
-- System mut: 'qso', if qso is not in available_mult, then this will be marked as "dup"
-- Expects { available_mults : { mults } , available_band_mult : { band_mode_mult_map } }
-- QSO will be a table of 3 fields: callsign, frequency, mode (mapped to the mode in available_modes)
function check_callsign(state, qso)
    return {
        available_mults = { "qso" },
        available_band_mult = {
            "160" = { "qso" },
            "80" = { "qso" },
            "40" = { "qso" },
            "20" = { "qso" },
            "15" = { "qso" },
            "10" = { "qso" },
        }
    }
end