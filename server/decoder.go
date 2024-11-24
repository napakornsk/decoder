package main

import (
	"context"
	"fmt"

	pb "decoder/proto"
)

// สัญลักษณ์ “L” หมายความว่า ตัวเลขด้านซ้าย มีค่ามากกว่า ตัวเลขด้านขวา
// สัญลักษณ์ “R” หมายความว่า ตัวเลขด้านขวา มีค่ามากกว่า ตัวเลขด้านซ้าย
// สัญลักษณ์ “=“ หมายความว่า ตัวเลขด้านซ้าย มีค่าเท่ากับ ตัวเลขด้านขวา

func (s *DecoderServiceServer) Decode(c context.Context, req *pb.DecoderRequest) (
	*pb.DecoderResponse, error) {
	fmt.Println("Decode was invoked")
	val := 0
	n := len(req.GetEncodeText())
	ans := make([]string, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			ans = append(ans, fmt.Sprint(val))
		} else {
			t := req.EncodeText[i-1]
			switch t {
			case 'L':
				l := val + 1
				r := val - 1
				if r < 0 {
					r = 0
				}
				ans[i-1] = fmt.Sprint(l)
				ans[i+1] = fmt.Sprint(r)
				ans[i] = fmt.Sprint(val)
			case 'R':
				l := val - 1
				r := val + 1
				if l < 0 {
					l = 0
				}
				ans[i-1] = fmt.Sprint(l)
				ans[i+1] = fmt.Sprint(r)
				ans[i] = fmt.Sprint(val)
			case '=':
				ans[i-1] = fmt.Sprint(val)
				ans[i+1] = fmt.Sprint(val)
				ans[i] = fmt.Sprint(val)
			default:
				ans[i] = fmt.Sprint(val)
			}
		}
	}

	final := ""
	for _, a := range ans {
		final += fmt.Sprint(a)
	}

	return &pb.DecoderResponse{
		Data: final,
	}, nil
}
