class ProductOfNumbers {
public:
    ProductOfNumbers():products_({1}) {
		
    }
    
    void add(int num) {
		if (num == 0) {
			products_ = {1};
		} else {
			products_.emplace_back(products_.back()*num);
		}
    }
    
    int getProduct(int k) {
		int i = products_.size() - k - 1;
		if (i < 0) {
			return 0;
		}
		return products_.back() / products_[i];
    }

private:
	vector<int> products_;
};

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * ProductOfNumbers* obj = new ProductOfNumbers();
 * obj->add(num);
 * int param_2 = obj->getProduct(k);
 */