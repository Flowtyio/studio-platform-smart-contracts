import DapperSport from "../../../contracts/DapperSport.cdc"

transaction(
    seriesID: UInt64,
    setID: UInt64,
    playID: UInt64,
    tier: String,
    maxMintSize: UInt64?,
   ) {
    // local variable for the admin reference
    let admin: &DapperSport.Admin

    prepare(signer: AuthAccount) {
        // borrow a reference to the Admin resource
        self.admin = signer.borrow<&DapperSport.Admin>(from: DapperSport.AdminStoragePath)
            ?? panic("Could not borrow a reference to the DapperSport Admin capability")
    }

    execute {
        let id = self.admin.createEdition(
            seriesID: seriesID,
            setID: setID,
            playID: playID,
            maxMintSize: maxMintSize,
            tier: tier,
        )

        log("====================================")
        log("New Edition:")
        log("EditionID: ".concat(id.toString()))
        log("====================================")
    }
}

