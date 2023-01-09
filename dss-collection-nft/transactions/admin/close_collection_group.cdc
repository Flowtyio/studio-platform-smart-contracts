import DSSCollection from 0xf8d6e0586b0a20c7
import NonFungibleToken from 0xf8d6e0586b0a20c7

transaction(collectionGroupID: UInt64) {
    // local variable for the admin reference
    let admin: &DSSCollection.Admin

    prepare(signer: AuthAccount) {
        // borrow a reference to the Admin resource
        self.admin = signer.borrow<&DSSCollection.Admin>(from: DSSCollection.AdminStoragePath)
            ?? panic("Could not borrow a reference to the DSSCollection Admin capability")
    }

    execute {
        let id = self.admin.closeCollectionGroup(
            id: collectionGroupID
        )

        log("====================================")
        log("Closed Collection Group:")
        log("CollectionGroupID: ".concat(id.toString()))
        log("====================================")
    }
}