// Copyright © 2017-2020 Trust Wallet.
//
// This file is part of Trust. The full Trust copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

#include <TrustWalletCore/TWAnyAddress.h>
#include "HexCoding.h"

#include "TestUtilities.h"
#include <gtest/gtest.h>

using namespace TW;

TEST(TWBinanceSmartChain, Address) {

    auto privateKey = WRAP(TWPrivateKey, TWPrivateKeyCreateWithData(DATA("727f677b390c151caf9c206fd77f77918f56904b5504243db9b21e51182c4c06").get()));
    auto publicKey = WRAP(TWPublicKey, TWPrivateKeyGetPublicKeySecp256k1(privateKey.get(), false));
    auto string = "0xf3d468DBb386aaD46E92FF222adDdf872C8CC064";
    
    auto address = WRAP(TWAnyAddress, TWAnyAddressCreateWithPublicKey(publicKey.get(), TWCoinTypeSmartChain));
    auto expected = WRAP(TWAnyAddress, TWAnyAddressCreateWithString(STRING(string).get(), TWCoinTypeSmartChain));

    auto addressString = WRAPS(TWAnyAddressDescription(address.get()));
    auto expectedString = WRAPS(TWAnyAddressDescription(expected.get()));

    assertStringsEqual(addressString, string);
    assertStringsEqual(expectedString, string);
}
