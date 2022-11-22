import { FC, useState } from 'react';
import { Button, Row, Col } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { Icon, BaseUserCard, DiffContent, FormatTime } from '@/components';
import { TIMELINE_NORMAL_ACTIVITY_TYPE } from '@/common/constants';
import * as Type from '@/common/interface';

const data1 = {
  title: '不是管理员，提一个问题看看能不能编辑resserved tag?',
  tags: [
    {
      display_name: 'bug',
      slug_name: 'bug',
      recommend: true,
      reserved: false,
    },
    {
      display_name: '黄马甲',
      slug_name: '黄马甲',
      recommend: false,
      reserved: true,
    },
    {
      display_name: 'go',
      slug_name: 'go',
      recommend: false,
      reserved: false,
    },
  ],
  content: `# 前言
  手写 Promise 是面试的时候大家都逃避的送命题，在学些了解后发现通过实现源码更能将新一代的异步方案理解的通透，知其然知其所以然的运用。

  如果直接将源码贴到此处势必不能有更大的收获，下面就按实现版本来看做简要分析。

  ## 回顾 Promise
  Promise 是 CommonJS 提出来的这一种规范，有多个版本，在 ES6 当中已经纳入规范，原生支持 Promise 对象，非 ES6 环境可以用类似 Bluebird、Q 这类库来支持。

  Promise 可以将回调变成链式调用写法，流程更加清晰，代码更加优雅，还可以批量处理异步任务。

  简单归纳下 Promise：三个状态、两个过程、一个方法，快速记忆方法：3-2-1

  三个状态：pending、fulfilled、rejected

  两个过程：

  * pending → fulfilled（resolve）
  * pending → rejected（reject）
  一个方法：then

  当然还有其他概念，如 catch、 Promise.all/race/allSettled。`,
};

const data2 = {
  title: '提一个问题看看能不能编辑 resserved tag?',
  tags: [
    {
      display_name: 'discussion',
      slug_name: 'discussion',
      recommend: true,
      reserved: false,
    },
    {
      display_name: '黄马甲',
      slug_name: '黄马甲',
      recommend: false,
      reserved: true,
    },
    {
      display_name: 'go',
      slug_name: 'go',
      recommend: false,
      reserved: false,
    },
  ],
  content: `# 前言
  手写 Promise 是面试的时候大家都逃避的送命题，在学些了解后发现通过实现源码更能将新一代的异步方案理解的通透，知用。

  ## 增加的titlte

  如果直接将源码贴到此处势必不能有更大的收获。

  ## 回顾 Promise
  Promise 是 CommonJS 规范，并且有多个版本，在 ES6 当中已经纳入规范，原生支持 Promise 对象，非 ES6 环境可以用类似 Bluebird、Q 这类库来支持。

  Promise 可以将回调变成链式调用写法，流程更加清晰，代码更加优雅，还可以批量处理异步任务。

  简单归纳下 Promise：三个状态、两个过程、一个方法，快速记忆方法：3-2-1

  两个过程：

  * pending → fulfilled（resolve）
  * pending → rejected（reject）
  一个方法：then
  `,
};

interface Props {
  data: Type.TimelineItem;
  objectInfo: Type.TimelineObject;
  source: 'question' | 'answer' | 'tag';
  isAdmin: boolean;
}
const Index: FC<Props> = ({
  data,
  isAdmin,
  source = 'question',
  objectInfo,
}) => {
  const { t } = useTranslation('translation', { keyPrefix: 'timeline' });
  const [isOpen, setIsOpen] = useState(false);
  const handleItemClick = () => {
    setIsOpen(!isOpen);
  };
  return (
    <>
      <tr>
        <td>
          <FormatTime time={data.created_at} />
          <br />
          {data.cancelled_at > 0 && <FormatTime time={data.cancelled_at} />}
        </td>
        <td>
          {(data.activity_type === 'rollback' ||
            data.activity_type === 'edited' ||
            data.activity_type === 'asked' ||
            data.activity_type === 'created' ||
            (source === 'answer' && data.activity_type === 'answered')) && (
            <Button
              onClick={handleItemClick}
              variant="link"
              className="text-body p-0 btn-no-border">
              <Icon
                name="caret-right-fill"
                className={`me-1 ${isOpen ? 'rotate-90-deg' : 'rotate-0-deg'}`}
              />
              {t(data.activity_type)}
            </Button>
          )}
          {data.activity_type === 'accept' && (
            <Link to={`/question/${objectInfo.question_id}`}>
              {t(data.activity_type)}
            </Link>
          )}

          {source === 'question' && data.activity_type === 'answered' && (
            <Link
              to={`/question/${objectInfo.question_id}/${objectInfo.answer_id}`}>
              {t(data.activity_type)}
            </Link>
          )}

          {data.activity_type === 'commented' && (
            <Link
              to={
                data.object_type === 'answer'
                  ? `/question/${objectInfo.question_id}/${objectInfo.answer_id}?commentId=${data.object_id}`
                  : `/question/${objectInfo.question_id}?commentId=${data.object_id}`
              }>
              {t(data.activity_type)}
            </Link>
          )}

          {TIMELINE_NORMAL_ACTIVITY_TYPE.includes(data.activity_type) && (
            <div>{t(data.activity_type)}</div>
          )}

          {data.cancelled && (
            <div className="text-danger"> {t('cancelled')}</div>
          )}
        </td>
        <td>
          {data.activity_type === 'downvote' && !isAdmin ? (
            <div>{t('n_or_a')}</div>
          ) : (
            <BaseUserCard
              className="fs-normal"
              data={{
                username: data.username,
                display_name: data.user_display_name,
              }}
              showAvatar={false}
              showReputation={false}
            />
          )}
        </td>
        <td>{data.comment}</td>
      </tr>
      <tr className={isOpen ? '' : 'd-none'}>
        {/* <td /> */}
        <td colSpan={5} className="p-0 py-5">
          <Row className="justify-content-center">
            <Col xxl={8}>
              <DiffContent currentData={data1} prevData={data2} />
            </Col>
          </Row>
        </td>
      </tr>
    </>
  );
};

export default Index;
