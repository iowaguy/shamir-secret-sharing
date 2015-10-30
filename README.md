# Shamir's Secret Sharing
A Go implementation of Adi Shamir's 1979 cryptosystem.

Read the original publication here: http://cs.jhu.edu/~sdoshi/crypto/papers/shamirturing.pdf

Or the nice Wikipedia summary here: https://en.wikipedia.org/wiki/Shamir's_Secret_Sharing

Server features are not functional yet.

To try it out, follow these steps:

# how to build it
# NOTE: must have your go workspace and GOPATH environment variable setup
# run the following command
go install github.com/iowaguy/shamir-secret-sharing/

# how to use it
# step 1: find executable
cd $GOPATH/bin

# step 2: make keys
./shamir-secret-sharing -k <type key here> <threshold #> <total keys to make>

# step 3: decode keys
./shamir-secret-sharing -d <key #1> <key #2> <key #3> <key #4> <...> <key #n>