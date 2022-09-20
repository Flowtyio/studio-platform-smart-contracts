import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import PackNFT from "../../contracts/PackNFT.cdc"

pub fun main(account: Address): [UInt64] {
    let receiver = getAccount(account)
        .getCapability(PackNFT.CollectionPublicPath)!
        .borrow<&{NonFungibleToken.CollectionPublic}>()!

    return receiver.getIDs()
}
