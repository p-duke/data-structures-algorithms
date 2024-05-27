require "./lib/binary_search/binary_search_practice"

describe BinarySearchPractice do
  describe "#search" do
    it "returns the index of the target if present" do
      nums = [-1,0,3,5,9,12]
      target = 9
      result = BinarySearchPractice.search(nums, target)

      expect(result).to eql(4)
    end

    it "returns -1 if the target isn't present" do
      nums = [-1,0,3,5,9,12]
      target = 2
      result = BinarySearchPractice.search(nums, target)

      expect(result).to eql(-1)
    end

    it "should work with single element arrays " do
      nums = [5]
      target = 5
      result = BinarySearchPractice.search(nums, target)

      expect(result).to eql(0)
    end
  end
end

