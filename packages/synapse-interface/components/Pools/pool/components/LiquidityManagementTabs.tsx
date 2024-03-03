import Tabs from '@tw/Tabs'
import TabItem from '@tw/TabItem'

const LiquidityManagementTabs = ({ cardNav, setCardNav }) => {
  return (
    <Tabs>
      <TabItem
        isActive={cardNav === 'addLiquidity'}
        onClick={() => {
          setCardNav('addLiquidity')
        }}
        className="rounded-tl-lg"
      >
        Add Liquidity
      </TabItem>
      <TabItem
        isActive={cardNav === 'removeLiquidity'}
        onClick={() => {
          setCardNav('removeLiquidity')
        }}
        className="rounded-tr-lg"
      >
        Remove Liquidity
      </TabItem>
    </Tabs>
  )
}
export default LiquidityManagementTabs