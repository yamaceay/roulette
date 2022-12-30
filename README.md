## Configuration of a Strategy
```
{
	"prob":     < success probability >,
	"wage":     < initial wage >,
	"stepFunc": < wage increasing strategy >,
	"stopLoss": < upper loss bound >,
	"winRound": < calculate until k-th round >,
    "meanShift": < expected winning probability of house [0 = perfectly fair, 1 = house always wins] >
}
```
<hr/>

### Stake Strategy

#### Start
_Explanation_: The game starts with an initial stake $f(0)$

#### Step
_Explanation_: Define a strategy for increasing the stakes, which considers two factors at hand:
* Recovery of Failed Bets: The overall outcome after losing in $X_1, …, X_{n-1}$ and winning in $X_n$.
* Final Potential Damage: The overall loss expected after losing in all $X_1, …, X_n$. 

_Options_:
* 'two' (Power of Two): After each losing, the stake is doubled. $f(k) = 2^k$
* 'fib' (Fibonacci): After losing k-1 bets, the stake is equal to the k-th Fibonacci number. $f(k) = fib(k)$.

_Default_: 'two' 

_Controls_: Financial Stability

#### Stop
_Explanation_: Under what circumstances the game should be stopped.

_Options_:
* 'stopLoss': If at least $N$ was lost ( = bankruptcy), stop the current increasing strategy.

<hr/>

### Simultaneous Games
_Explanation_: If $n$ games are played at once, this leads to a standard normal distribution: $Y = \underset{i \in [n]}\sum \frac{X_i - \mu}{\sigma \sqrt{n}}$ with $\mu = -\frac{1}{74} \approx -0.014$ and $\sigma = \frac{1341}{1369} \approx 0.989$

_Critique_:

On the one hand: The more games are played simultaneously, the more likely the expected outcome is negative, so the number of games should not be very high.

On the other hand: Rare events can affect the whole system, if the number of games are very low, so the number of games should not be very low.

_Controls_: Stability of Results

<hr/>

### Success Probability [EXAMPLES]
_Explanation_: How high the probability of success is for each round.

_Options_: 
* 'xl-risky': $p = 0.1$
* 'l-risky': $p = 0.2$
* 'm-risky': $p = 0.3$
* 's-risky': $p = 0.4$
* 'fair': $p = 0.5$
* 's-safe': $p = 0.6$
* 'm-safe': $p = 0.7$
* 'l-safe': $p = 0.8$
* 'xl-safe': $p = 0.9$

_Default_: 'fair'

_Controls_: Winning probability

<hr/>

### Dependence on Prior Results [IT DOESN'T MATTER]
_Explanation_: Given $X_1, …, X_{n-1}$ events with identical outcomes, predict $X_n$.

_Options_:
* 'corr' (Correlation): Choose $X_n$ the same as before
* 'regr' (Regression to the Mean): Pick the opposite for $X_n$

_Default_: 'regr'

_Disclaim_: Both assumptions are false. There is no proven hidden connection between a current event and events occurred previously. But one of the both assumptions appear to be more reliable in various cases.

_Controls_: Playing strategy

<hr/>
