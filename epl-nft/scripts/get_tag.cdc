import EnglishPremierLeague from "./EnglishPremierLeague.cdc"

pub fun main(tagID: UInt64): EnglishPremierLeague.Tag {
    return EnglishPremierLeague.getTag(id: tagID)!
}