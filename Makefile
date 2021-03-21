 # 
 .PHONY: all
 all: $(PWD)/pkg/tokens/trontoken/trontoken.go

$(PWD)/pkg/tokens/trontoken/trontoken.go: $(PWD)/solidity/trontoken.sol
    # STEP 1: Create a docker image of the abigen tool 
	$(PWD)/tools/ops.sh image 0.4.26
	# STEP 2: Run the abigen image to process the solidity contract
	docker run --rm -v $(PWD)/solidity/trontoken.sol:/opt/solidity/trontoken.sol \
	                -v $(PWD)/pkg/tokens/trontoken/:/opt/trontoken  \
					oc/abigentool:0.4.26 --sol /opt/solidity/trontoken.sol \
					                 	 --pkg trontoken \
									 	 --out /opt/trontoken/trontoken.go

clean:
	rm -rf $(PWD)/pkg/*

	