--
-- Note: The contest descriptor should not use any global variables.
-- Functions will be provided with a "state" for saving the state
--

--
-- MCL runtime provides the following lookup function:
-- callsign_cq_zone()   -- Return the inferred CQ zone of the callsign, 0 for unknown
-- callsign_itu_zone()  -- Return the inferred ITU zone of the callsign, 0 for unknown
-- callsign_continent() -- Return the inferred continent of the callsign, "" for unknown
-- callsign_dxcc()      -- Return the inferred DXCC of the callsign, 0 for unknown
--

-- Basic functions for getting contest metadata
function LoadMetadata(meta)
    -- basic metadata of the contest
    -- meta_version: the schema version of the metadata
    meta.ApiVersion = "1"
    -- version: the version of the contest, for identifying other hosts on the network and display.
    meta.Version = "2022.01"

    -- the name of the contest, for identifying other hosts on the network and display.
    meta.ContestName = "CQWW SSB"

    -- the modes of the contest
    meta.AvailableModes = { "SSB" }

    -- the bands of the contest
    meta.AvailableBands = { "160", "80", "40", "20", "15", "10" }

    -- the sent exchange data of the contest
    meta.ExchangeSent = { "rst_sent", "zone_sent" }
    -- the received exchange data of the contest
    meta.ExchangeRcvd = { "rst_rcvd", "zone_rcvd" }

    -- multipliers are only used for display, score calculation is fully handled in the script
    meta.Multipliers = { "Zone" }

    meta.FieldInfos = {
        rst_sent = {
            DisplayName = "Tx RST",
            Description = "RST Sent",
            FieldType = "tx",
            AdifName = "RST_SENT",
            ValueRegex = "^[1-5][1-9][1-9]$",
        },
        rst_rcvd = {
            DisplayName = "Rx RST",
            Description = "RST Received",
            FieldType = "rx",
            AdifName = "RST_RCVD",
            ValueRegex = "^[1-5][1-9][1-9]$",
        },
        zone_sent = {
            DisplayName = "TxZone",
            Description = "CQ Zone Sent",
            FieldType = "tx_const",
            AdifName = "MY_CQ_ZONE",
            ValueRegex = "^\\d{1,2}$",
        },
        zone_rcvd = {
            DisplayName = "RxZone",
            Description = "CQ Zone Received",
            FieldType = "rx",
            AdifName = "CQZ",
            ValueRegex = "^\\d{1,2}$",
        },
        continent_ = {
            DisplayName = "Continent",
            Description = "Continent",
            FieldType = "info",
            AdifName = "",
            ValueRegex = "",
            ValidValues = { "N/A", "AF", "AN", "AS", "EU", "NA", "OC", "SA" },
        }
    }
end

function DraftQSO(qso)
    qso.ExchangeSent["rst_sent"] = "599"
    qso.ExchangeRcvd["rst_rcvd"] = "599"
    qso.Expect = "zone_rcvd"
    return true
end

-- Create a new state of the contest
-- MCL will not inteprete the content of the state, it is totally private.
function init_state(station)
    return {
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
        title = { "Band", "QSO", "Zone", "Score", "Score/Q" },
        fields = {
            { "80", 0, 0, 0, 0 },
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
    }
end
