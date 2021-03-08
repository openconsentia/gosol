 
 .PHONY: all
 all: $(PWD)/pkg/trontoken/trontoken.go

$(PWD)/pkg/trontoken/trontoken.go: $(PWD)/solidity/trontoken.sol
	$(PWD)/tools/ops.sh image 0.4.26
	docker run --rm -v $(PWD)/solidity/trontoken.sol:/opt/solidity/trontoken.sol \
	                -v $(PWD)/pkg/trontoken/:/opt/trontoken  \
					oc/abigentool:0.4.26 --sol /opt/solidity/trontoken.sol \
					                 	 --pkg trontoken \
									 	 --out /opt/trontoken/trontoken.go

clean:
	rm -rf $(PWD)/pkg/*

	