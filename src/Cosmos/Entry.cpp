// Copyright © 2017-2022 Trust Wallet.
//
// This file is part of Trust. The full Trust copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

#include "Entry.h"

#include "Address.h"
#include "Signer.h"

using namespace TW;
using namespace std;

namespace TW::Cosmos {

// Note: avoid business logic from here, rather just call into classes like Address, Signer, etc.

bool Entry::validateAddress(TWCoinType coin, const string& address, TW::byte, TW::byte, const char* hrp) const {
    if (hrpForString(hrp) != TWHRPUnknown) {
        return Address::isValid(coin, address);
    }
    return Address::isValid(address, hrp);
}

string Entry::deriveAddress(TWCoinType coin, const PublicKey& publicKey, TW::byte, const char* hrp) const {
    if (!std::string(hrp).empty()) {
        return Address(hrp, publicKey, coin).string();
    }
    return Address(coin, publicKey).string();
}

Data Entry::addressToData([[maybe_unused]] TWCoinType coin, const std::string& address) const {
    Address addr;
    if (!Address::decode(address, addr)) {
        return Data();
    }
    return addr.getKeyHash();
}

void Entry::sign(TWCoinType coin, const TW::Data& dataIn, TW::Data& dataOut) const {
    auto input = Proto::SigningInput();
    input.ParseFromArray(dataIn.data(), (int)dataIn.size());
    auto serializedOut = Signer::sign(input, coin).SerializeAsString();
    dataOut.insert(dataOut.end(), serializedOut.begin(), serializedOut.end());
}

string Entry::signJSON(TWCoinType coin, const std::string& json, const Data& key) const {
    return Signer::signJSON(json, key, coin);
}

} // namespace TW::Cosmos
